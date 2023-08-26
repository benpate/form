package form

import (
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
		return nil, derp.NewInternalError("form.Widget", "Unrecognized form widget", element)
	}

	return widget, nil
}

func (element *Element) View(schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	widget, err := element.Widget()

	if err != nil {
		return err
	}

	return widget.View(element, schema, lookupProvider, value, b)

}

func (element *Element) Edit(schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	widget, err := element.Widget()

	if err != nil {
		return err
	}

	return widget.Edit(element, schema, lookupProvider, value, b)

}

// GetValue returns the value of the element at the provided path.  If the schema is present,
// then it is used to resolve the value.  If the schema is not present, then the value is returned using path lookup instead.
func (element *Element) GetString(value any, s *schema.Schema) string {
	return convert.String(element.getValue(value, s))
}

func (element *Element) GetSliceOfString(value any, s *schema.Schema) []string {
	return convert.SliceOfString(element.getValue(value, s))
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

func (element *Element) getElement(s *schema.Schema) schema.Element {

	// If there is a schema, use it to get the value
	if s != nil {
		schemaElement, _ := s.GetElement(element.Path)
		return schemaElement
	}

	return nil
}

func (element *Element) inputVisible(s *schema.Schema, values any) bool {

	// RULE: if the element is read-only, then it should not be used as an input value
	if element.ReadOnly {
		return false
	}

	// If a "show-if" option is present, then we need to evaluate it to see if the input value should be considered "present" or not.
	if showIf := element.Options.GetString("show-if"); showIf != "" {
		return s.Match(values, exp.Parse(showIf))
	}

	return true
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

/******************************
 * SERIALIZATION METHODS
 ******************************/

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
				return derp.NewInternalError("form.UnmarshalMap", "Error parsing child form information.", childInterface)
			}
		}
	}

	return nil
}
