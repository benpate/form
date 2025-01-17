package form

import (
	"net/url"

	"github.com/benpate/derp"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/rs/zerolog/log"
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
func (form *Form) SetURLValues(object any, value url.Values, lookupProvider LookupProvider) error {

	const location = "form.Form.SetFromURLValues"

	// Try to apply all values from the form to the object
	for _, element := range form.Element.AllElements() {

		// RULE: do not update fields that are not visible
		if !element.isInputVisible(&form.Schema, value) {
			continue
		}

		// Try to replace new lookup codes (if neede)
		newValue, updated, err := element.replaceNewLookup(lookupProvider, value.Get(element.Path))

		if err != nil {
			return derp.Wrap(err, location, "Error writing new lookup value")
		}

		if updated {
			value[element.Path] = []string{newValue}
		}

		// Update the original object with the new value
		// Errors are intentionally ignored here.
		// Unallowed data does not make it through the schema filter
		// nolint: errcheck
		if err := form.Schema.Set(object, element.Path, value[element.Path]); err != nil {
			log.Debug().Err(err).Str("path", element.Path).Msg("Error setting value")
		}
	}

	// Validate that all of the data in the object are valid.
	if err := form.Schema.Validate(object); err != nil {
		return derp.Wrap(err, location, "Error validating object")
	}

	return nil
}

// Do applies all of the data from the value map into the target object
func (form *Form) SetAll(object any, value mapof.Any, lookupProvider LookupProvider) error {

	const location = "form.Form.SetAll"

	// Try to apply all values from the form to the object
	for _, element := range form.Element.AllElements() {

		// RULE: Do not update invisible fields
		if !element.isInputVisible(&form.Schema, value) {
			continue
		}

		// Try to replace new lookup codes (if needed)
		newValue, updated, err := element.replaceNewLookup(lookupProvider, value.GetString(element.Path))

		if err != nil {
			return derp.Wrap(err, location, "Error writing new lookup value")
		}

		if updated {
			value.SetString(element.Path, newValue)
		}

		// Update the original object with the new value
		if err := form.Schema.Set(object, element.Path, value[element.Path]); err != nil {
			return derp.Wrap(err, location, "Error setting value", element.Path)
		}
	}

	// Validate that all of the data in the object are valid.
	if err := form.Schema.Validate(object); err != nil {
		return derp.Wrap(err, location, "Error validating object")
	}

	return nil
}

// Encoding returns the "enctype" attribute for the form.
// Default is ""
func (form *Form) Encoding() string {
	return form.Element.Encoding()
}
