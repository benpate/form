package form

import (
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/exp"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
)

// Element defines a single form element, or a nested form layout.  It can be serialized to and from a database.
type Element struct {
	Type        string    `json:"type"`                  // The kind of form element
	ID          string    `json:"id"`                    // The ID of the element (needed by some widgets)
	Path        string    `json:"path"`                  // Path to the data value displayed in for this form element
	Label       string    `json:"label,omitempty"`       // Short label to be displayed on the form element
	Description string    `json:"description,omitempty"` // Longer description text to be displayed on the form element
	Options     mapof.Any `json:"options,omitempty"`     // Additional custom properties defined by individual widgets
	Children    []Element `json:"children,omitempty"`    // Array of sub-form elements that may be displayed depending on the kind.
	ReadOnly    bool      `json:"readOnly,omitempty"`    // If true, then this element is read-only
}

func NewElement() Element {
	return Element{
		Options:  make(mapof.Any),
		Children: make([]Element, 0),
	}
}

func (element *Element) Widget() (Widget, error) {

	widget, ok := registry[element.Type]

	if !ok {
		return nil, derp.InternalError("form.Widget", "Unrecognized form widget", element)
	}

	return widget, nil
}

func (element *Element) View(form *Form, lookupProvider LookupProvider, value any, b *html.Builder) error {

	widget, err := element.Widget()

	if err != nil {
		return err
	}

	return widget.View(form, element, lookupProvider, value, b)

}

func (element *Element) Edit(form *Form, lookupProvider LookupProvider, value any, b *html.Builder) error {

	widget, err := element.Widget()

	if err != nil {
		return err
	}

	return widget.Edit(form, element, lookupProvider, value, b)
}

func (element *Element) Encoding() string {

	if widget, err := element.Widget(); err == nil {
		return widget.Encoding(element)
	}

	return ""
}

// GetValue returns the value of the element at the provided path.  If the schema is present,
// then it is used to resolve the value.  If the schema is not present, then the value is returned using path lookup instead.
func (element *Element) GetString(value any, s *schema.Schema) string {
	return convert.String(element.getValue(value, s))
}

// GetSliceOfString rturns a slice of strings for a provided path.
func (element *Element) GetSliceOfString(value any, s *schema.Schema) []string {
	return convert.SliceOfString(element.getValue(value, s))
}

// IsEmpty returns TRUE if the element is not defined
func (element Element) IsEmpty() bool {
	return element.Type == ""
}

// GetSchema finds and returns the schema.Element associated with this Element path
func (element *Element) GetSchema(s *schema.Schema) schema.Element {

	// If there is a schema, use it to get the value
	if s != nil {
		schemaElement, _ := s.GetElement(element.Path)
		return schemaElement
	}

	return nil
}

// getValue returns the value of the element at the provided path.  If the schema is present,
// then it is used to resolve the value.  If the schema is not present, then the value is returned using path lookup instead.
func (element *Element) getValue(value any, s *schema.Schema) any {

	// If there is a schema, use it to get the value
	if s != nil {
		result, _ := s.Get(value, element.Path)
		return result
	}

	return nil
}

func (element *Element) isInputVisible(s *schema.Schema, value any) (bool, error) {

	// RULE: ReadOnly elements are not "visible" in the form
	if element.ReadOnly {
		return false, nil
	}

	// Collect the "show-if" property of the form element
	showIf := element.Options.GetString("show-if")

	// RULE: if there is no "show-if" option, then the input is always visible
	if showIf == "" {
		return true, nil
	}

	// Parse the "show-if" text into a valid expression
	expression := exp.Parse(showIf)

	// If the data matches the expression then the input is visible
	visible, err := s.Match(value, expression)

	if err != nil {
		return false, derp.Wrap(err, "form.element.isInputVisible", "Error evaluating show-if expression", showIf)
	}

	// Success
	return visible, nil
}

// replaceLookupValue
func (element Element) replaceNewLookup(lookupProvider LookupProvider, value string) (string, bool, error) {

	const location = "form.element.replaceNewLookup"

	// RULE: lookupProvider must not be nil
	if lookupProvider == nil {
		return value, false, nil
	}

	// RULE: value MUST match the ::NEWVALUE:: new item identifier
	if !strings.HasPrefix(value, NewItemIdentifier) {
		return value, false, nil
	}

	// Get the lookup provider name
	groupName := element.Options.GetString("provider")

	if groupName == "" {
		return value, false, nil
	}

	// Use the LookupProvider to get the LookupGroup
	lookupGroup := lookupProvider.Group(groupName)

	if lookupGroup == nil {
		return value, false, nil
	}

	// If the LookupGroup is writable, then try to add a new value
	if writableGroup, ok := lookupGroup.(WritableLookupGroup); ok {

		value = strings.TrimPrefix(value, NewItemIdentifier)
		value, err := writableGroup.Add(value)

		if err != nil {
			return value, false, derp.Wrap(err, location, "Error adding new lookup value", groupName, element.Path, value)
		}

		return value, true, nil
	}

	// Otherwise, the value cannot be written, so keep the original value
	return value, false, nil
}

// 	Autocomplete string `json:"autocomplete"` // https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/autocomplete

// AllPaths returns pointers to all of the valid paths in this form
func (element *Element) AllElements() []*Element {

	var result []*Element

	if element.ReadOnly {
		return result
	}

	// If THIS element has a Path, then add it to the result
	if element.Path == "" {
		result = []*Element{}
	} else {
		result = []*Element{element}
	}

	// Scan all chiild elements for THEIR paths, and add them to the result
	for index := range element.Children {
		result = append(result, element.Children[index].AllElements()...)
	}

	// Success
	return result
}

/***********************************
 * Serialization Methods
 ***********************************/

// UnmarshalMap parses data from a generic structure (mapof.Any) into a Form record.
func (element *Element) UnmarshalMap(data map[string]any) error {

	element.Type = convert.String(data["type"])
	element.Path = convert.String(data["path"])
	element.Label = convert.String(data["label"])
	element.Description = convert.String(data["description"])
	element.ReadOnly = convert.Bool(data["readOnly"])

	element.Options = make(mapof.Any)
	if options, ok := data["options"].(map[string]any); ok {
		element.Options = options
	}

	if children, ok := data["children"].([]any); ok {
		element.Children = make([]Element, len(children))
		for index, childInterface := range children {
			if childData, ok := childInterface.(map[string]any); ok {
				var child Element
				if err := child.UnmarshalMap(childData); err != nil {
					return derp.Wrap(err, "form.UnmarshalMap", "Error parsing child form information.", childInterface)
				}
				element.Children[index] = child
			} else {
				return derp.InternalError("form.UnmarshalMap", "Error parsing child form information.", childInterface)
			}
		}
	}

	return nil
}
