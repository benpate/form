package form

import (
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
)

type Form struct {
	Schema  schema.Schema
	Element Element
}

// New returns a fully initialized Form object (with all required values)
func New(schema schema.Schema, element Element) Form {
	return Form{
		Schema:  schema,
		Element: element,
	}
}

// Viewer creates an in-place form and executes its "Viewer" method
func Viewer(schema schema.Schema, element Element, value any, lookupProvider LookupProvider) (string, error) {
	form := New(schema, element)
	return form.Viewer(value, lookupProvider)
}

// Viewer creates an in-place form and executes its "Editorr" method
func Editor(schema schema.Schema, element Element, value any, lookupProvider LookupProvider) (string, error) {
	form := New(schema, element)
	return form.Editor(value, lookupProvider)
}

/********************************
 * Drawing Methods
 ********************************/

// DrawString() generates this form as a string
func (form *Form) Editor(value any, lookupProvider LookupProvider) (string, error) {
	builder := html.New()
	err := form.BuildEditor(value, lookupProvider, builder)
	return builder.String(), err
}

// DrawString() generates this form as a string
func (form *Form) Viewer(value any, lookupProvider LookupProvider) (string, error) {
	builder := html.New()
	err := form.BuildViewer(value, lookupProvider, builder)
	return builder.String(), err
}

// BuildEditor generates an editable view of this form
func (form *Form) BuildEditor(value any, lookupProvider LookupProvider, builder *html.Builder) error {
	return form.Element.Edit(&form.Schema, lookupProvider, value, builder)
}

// BuildViewer generates a read-only view of this form
func (form *Form) BuildViewer(value any, lookupProvider LookupProvider, builder *html.Builder) error {
	return form.Element.View(&form.Schema, lookupProvider, value, builder)
}

/********************************
 * Data Update Methods
 ********************************/

// Do applies all of the data from the value map into the target object
func (form *Form) SetAll(object any, value mapof.Any, lookupProvider LookupProvider) error {

	const location = "form.Form.SetAll"

	// Replace "NEW" values in LookupCodes
	if err := form.replaceNewLookups(value, lookupProvider); err != nil {
		return derp.Wrap(err, location, "Error replacing new lookups")
	}

	// Try to apply all values from the form to the object
	for _, element := range form.Element.AllElements() {

		if elementValue, ok := value[element.Path]; ok {
			if element.isInputVisible(&form.Schema, value) {
				if err := form.Schema.Set(object, element.Path, elementValue); err != nil {
					return derp.Wrap(err, location, "Error setting value", element.Path, elementValue)
				}
			}
		}
	}

	// Validate that all of the data in the object are valid.
	if err := form.Schema.Validate(object); err != nil {
		return derp.Wrap(err, location, "Error validating object")
	}

	return nil
}

func (form *Form) replaceNewLookups(value mapof.Any, lookupProvider LookupProvider) error {

	const newItemIdentifier = "::NEWVALUE::"

	if lookupProvider == nil {
		return nil
	}

	for _, element := range form.Element.AllElements() {

		// Get the original form value
		formValue, ok := value.GetStringOK(element.Path)

		if !ok {
			continue
		}

		// Value MUST match the "new item" identifier
		if !strings.HasPrefix(formValue, newItemIdentifier) {
			continue
		}

		// Get the lookup provider name
		providerName, ok := element.Options.GetStringOK("provider")

		if !ok {
			continue
		}

		if providerName == "" {
			continue
		}

		// User the LookupProvider to get the ProviderGroup
		group := lookupProvider.Group(providerName)

		if group == nil {
			continue
		}

		// If the group is writable, then try to add a new value
		if writableGroup, ok := group.(WritableLookupGroup); ok {

			formValue = strings.TrimPrefix(formValue, newItemIdentifier)
			formValue, err := writableGroup.Add(formValue)

			if err != nil {
				return derp.Wrap(err, "form.Form.replaceNewLookups", "Error adding new lookup value", element.Path, formValue)
			}

			value.SetString(element.Path, formValue)
		}
	}

	// Woot.
	return nil
}

// Encoding returns the "enctype" attribute for the form.
// Default is ""
func (form *Form) Encoding() string {
	return form.Element.Encoding()
}
