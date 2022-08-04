package form

import (
	"github.com/benpate/derp"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/maps"
	"github.com/benpate/rosetta/path"
	"github.com/benpate/rosetta/schema"
)

// Element defines a single form element, or a nested form layout.  It can be serialized to and from a database.
type Element struct {
	Type        string    `json:"type"`                  // The kind of form element
	ID          string    `json:"id"`                    // The ID of the element (needed by some widgets)
	Path        string    `json:"path"`                  // Path to the data value displayed in for this form element
	Label       string    `json:"label,omitempty"`       // Short label to be displayed on the form element
	Description string    `json:"description,omitempty"` // Longer description text to be displayed on the form element
	Options     maps.Map  `json:"options,omitempty"`     // Additional custom properties defined by individual widgets
	Children    []Element `json:"children,omitempty"`    // Array of sub-form elements that may be displayed depending on the kind.
}

func NewElement() Element {
	return Element{
		Options:  make(map[string]any),
		Children: make([]Element, 0),
	}
}

// HTML returns a populated HTML string for the provided value
func (element *Element) HTML(value any, schema *schema.Schema, lookupProvider LookupProvider) (string, error) {

	b := html.New()

	if err := element.WriteHTML(schema, lookupProvider, value, b); err != nil {
		return "", derp.Wrap(err, "form.HTML", "Error rendering element", element)
	}

	return b.String(), nil
}

func (element *Element) WriteHTML(schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	widgetFunc, ok := registry[element.Type]

	if !ok {
		return derp.NewInternalError("form.html", "Unrecognized form widget", element)
	}

	return widgetFunc(element, schema, lookupProvider, value, b)
}

// GetValue returns the value of the element at the provided path.  If the schema is present,
// then it is used to resolve the value.  If the schema is not present, then the value is returned using path lookup instead.
func (element *Element) GetString(value any, s *schema.Schema) (string, schema.Element) {

	result, schemaElement := element.getValue(value, s)

	switch schemaElement := schemaElement.(type) {
	case schema.Array:
		return convert.JoinString(result, schemaElement.Delimiter), schemaElement.Items
	}

	return convert.String(result), schemaElement
}

func (element *Element) GetSliceOfString(value any, s *schema.Schema) ([]string, schema.Element) {

	result, schemaElement := element.getValue(value, s)

	switch schemaElement := schemaElement.(type) {
	case schema.Array:
		return convert.SplitSliceOfString(result, schemaElement.Delimiter), schemaElement.Items
	}
	return convert.SliceOfString(result), schemaElement
}

// getValue returns the value of the element at the provided path.  If the schema is present,
// then it is used to resolve the value.  If the schema is not present, then the value is returned using path lookup instead.
func (element *Element) getValue(value any, s *schema.Schema) (any, schema.Element) {

	// If there is a schema, use it to get the value
	if s != nil {
		result, schemaElement, _ := s.Get(value, element.Path)
		return result, schemaElement
	}

	// Fall back to using path lookup.
	return path.Get(value, element.Path), nil
}

// 	Autocomplete string `json:"autocomplete"` // https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/autocomplete

// AllPaths returns pointers to all of the valid paths in this form
func (element *Element) AllElements() []*Element {

	var result []*Element

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

// UnmarshalMap parses data from a generic structure (map[string]any) into a Form record.
func (element *Element) UnmarshalMap(data map[string]any) error {

	element.Type = convert.String(data["type"])
	element.Path = convert.String(data["path"])
	element.Label = convert.String(data["label"])
	element.Description = convert.String(data["description"])

	element.Options = make(map[string]any)
	if options, ok := data["options"].(map[string]any); ok {
		for key, value := range options {
			element.Options[key] = convert.String(value)
		}
	}

	if children, ok := data["children"].([]any); ok {
		element.Children = make([]Element, len(children))
		for index, childInterface := range children {
			if childData, ok := childInterface.(map[string]any); ok {
				var child Element
				child.UnmarshalMap(childData)
				element.Children[index] = child
			} else {
				return derp.NewInternalError("form.UnmarshalMap", "Error parsing child form information.", childInterface)
			}
		}
	}

	return nil
}
