package form

import (
	"net/url"
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/sliceof"
	"github.com/rs/zerolog/log"
)

// Form defines all of the data for this JSON form.  This can be marshalled/unmarshalled
// as JSON, or entered directly into Go source code.
type Form struct {
	Schema  schema.Schema `json:"schema"`
	Element Element       `json:"form"`
	Options []string      `json:"options,omitempty"`
}

// New returns a fully initialized Form object (with all required values)
func New(schema schema.Schema, element Element, options ...string) Form {
	return Form{
		Schema:  schema,
		Element: element,
		Options: options,
	}
}

// Viewer creates an in-place form and executes its "Viewer" method
func Viewer(schema schema.Schema, element Element, value any, lookupProvider LookupProvider) (string, error) {
	form := New(schema, element)
	return form.Viewer(value, lookupProvider)
}

// Editor creates an in-place form and executes its "Editor" method
func Editor(schema schema.Schema, element Element, value any, lookupProvider LookupProvider, options ...string) (string, error) {
	form := New(schema, element, options...)
	return form.Editor(value, lookupProvider)
}

/***********************************
 * Drawing Methods
 ***********************************/

// Editor returns a fully populated HTML form for this Form
func (form *Form) Editor(value any, lookupProvider LookupProvider) (string, error) {
	builder := html.New()
	err := form.BuildEditor(value, lookupProvider, builder)
	return builder.String(), err
}

// Viewer returns a read-only HTML representation of this Form
func (form *Form) Viewer(value any, lookupProvider LookupProvider) (string, error) {
	builder := html.New()
	err := form.BuildViewer(value, lookupProvider, builder)
	return builder.String(), err
}

// BuildEditor generates an editable view of this form
func (form *Form) BuildEditor(value any, lookupProvider LookupProvider, builder *html.Builder) error {
	return form.Element.Edit(form, lookupProvider, value, builder)
}

// BuildViewer generates a read-only view of this form
func (form *Form) BuildViewer(value any, lookupProvider LookupProvider, builder *html.Builder) error {
	return form.Element.View(form, lookupProvider, value, builder)
}

/********************************
 * Data Update Methods
 ********************************/

// SetURLValues applies all of the data from the value map into the target object
func (form *Form) SetURLValues(object any, values url.Values, lookupProvider LookupProvider) error {

	const location = "form.Form.SetURLValues"

	// First, scan elements WITHOUT a "show-if" attribute.
	// Second, scan elements WITH a "show-if" attribute.
	// We do this so that dependent fields are calculated AFTER the parent fields are set.
	for _, showIf := range []bool{false, true} {

		for _, element := range form.Element.AllElements() {

			// RULE: Never update read-only fields
			if element.ReadOnly {
				continue
			}

			// Does this element have a "show-if" attribute? And, does it match the current scan?
			if hasShowIf := element.Options.GetString("show-if") != ""; hasShowIf != showIf {
				continue
			}

			// RULE: do not update fields that are not visible
			visible, err := element.isInputVisible(&form.Schema, object)

			if err != nil {
				return derp.Wrap(err, location, "Unable to evaluate show-if expression", element.Options.GetString("show-if"))
			}

			if !visible {
				continue
			}

			// Try to replace new lookup codes (if needed)
			newValue, updated, err := element.replaceNewLookup(lookupProvider, values.Get(element.Path))

			if err != nil {
				return derp.Wrap(err, location, "Unable to write new lookup value")
			}

			if updated {
				values[element.Path] = []string{newValue}
			}

			// Get the Widget associated with this Element
			widget, err := element.Widget()

			if err != nil {
				return derp.Wrap(err, location, "Unable to locate widget for element", element)
			}

			// If this element has a custom `SetURLValues` function, then
			// use that instead of the default value
			if setter, isSetter := widget.(URLValueSetter); isSetter {

				if err := setter.SetURLValue(form, element, object, values); err != nil {
					log.Debug().Err(err).Str("path", element.Path).Msg("Unable to set form value")
				}

				continue
			}

			// Update the original object with the new value. A rejected value is logged
			// rather than returned, so that one bad field does not abandon the rest of
			// the form; the schema has already filtered out anything it does not allow.
			if err := form.Schema.Set(object, element.Path, form.schemaSafeValue(element.Path, values)); err != nil {
				log.Debug().Err(err).Str("path", element.Path).Msg("Unable to set value")
			}
		}
	}

	// Success
	return nil
}

// schemaSafeValue shapes a posted url.Values entry for schema.Set.  url.Values holds a
// plain []string, but rosetta validates Array schemas through its ArrayGetterSetter
// interface, which the builtin slice does not implement (multi-value widgets like
// multiselect and check-button-group post these).  Wrapping Array-typed paths in a
// *sliceof.String bridges the two; all other paths pass through unchanged.  A path with
// no posted values wraps an empty slice, so un-checking every option clears the array.
func (form *Form) schemaSafeValue(path string, values url.Values) any {

	if element, ok := form.Schema.GetElement(path); ok {
		if _, isArray := element.(schema.Array); isArray {
			result := sliceof.String(values[path])
			return &result
		}
	}

	return values[path]
}

// Encoding returns the "enctype" attribute for the form.
// Default is ""
func (form *Form) Encoding() string {
	return form.Element.Encoding()
}

// OptionString returns the string value of a Form option.
func (form *Form) OptionString(name string) string {

	for _, option := range form.Options {
		if strings.HasPrefix(option, name+":") {
			return strings.TrimPrefix(option, name+":")
		}
	}
	return ""
}

// OptionInt returns the integer value of a Form option.
func (form *Form) OptionInt(name string) int {

	for _, option := range form.Options {
		if strings.HasPrefix(option, name+":") {
			optionString := strings.TrimPrefix(option, name+":")
			optionString = strings.TrimSpace(optionString)
			return convert.Int(optionString)
		}
	}

	return 0
}

// OptionBool returns the boolean value of a Form option.
func (form *Form) OptionBool(name string) bool {

	for _, option := range form.Options {
		if strings.HasPrefix(option, name+":") {
			optionString := strings.TrimPrefix(option, name+":")
			optionString = strings.TrimSpace(optionString)
			return convert.Bool(optionString)
		}
	}

	return false
}
