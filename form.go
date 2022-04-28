package form

import (
	"encoding/json"

	"github.com/benpate/convert"
	"github.com/benpate/datatype"
	"github.com/benpate/derp"
	"github.com/benpate/html"
	"github.com/benpate/schema"
)

// Form defines a single form element, or a nested form layout.  It can be serialized to and from a database.
type Form struct {
	Path        string `json:"path"`                  // Path to the data value displayed in for this form element
	Kind        string `json:"kind"`                  // The kind of form element
	ID          string `json:"id,omitempty"`          // DOM ID to use for this element.
	Label       string `json:"label,omitempty"`       // Short label to be displayed on the form element
	Description string `json:"description,omitempty"` // Longer description text to be displayed on the form element
	CSSClass    string `json:"cssClass,omitempty"`    // CSS Class override to apply to this widget.  This should be used sparingly
	Options     Map    `json:"options,omitempty"`     // Additional custom properties defined by individual widgets
	Children    []Form `json:"children,omitempty"`    // Array of sub-form elements that may be displayed depending on the kind.
	Show        Rule   `json:"show"`                  // Rules for showing/hiding/disabling this element
}

type Map map[string]interface{}

// NewForm returns a fully initialized Form object
func NewForm(kind string) Form {
	return Form{
		Kind:     kind,
		Options:  make(Map),
		Children: make([]Form, 0),
	}
}

// Parse attempts to convert interface{} value into a Form.
func Parse(data interface{}) (Form, error) {

	var result Form

	switch data := data.(type) {

	case datatype.Map:
		err := result.UnmarshalMap(map[string]interface{}(data))
		return result, err

	case map[string]interface{}:
		err := result.UnmarshalMap(data)
		return result, err

	case []byte:
		err := json.Unmarshal(data, &result)
		return result, err

	case string:
		err := json.Unmarshal([]byte(data), &result)
		return result, err

	}

	return result, derp.NewInternalError("form.Parse", "Cannot Parse Value: Unknown Datatype", data)
}

// MustParse guarantees that a value has been parsed into a Form, or else it panics the application.
func MustParse(data interface{}) Form {

	result, err := Parse(data)

	if err != nil {
		panic(err)
	}

	return result
}

// UnmarshalMap parses data from a generic structure (map[string]interface{}) into a Form record.
func (form *Form) UnmarshalMap(data map[string]interface{}) error {

	form.Path = convert.String(data["path"])
	form.Kind = convert.String(data["kind"])
	form.ID = convert.String(data["id"])
	form.Label = convert.String(data["label"])
	form.Description = convert.String(data["description"])
	form.CSSClass = convert.String(data["cssClass"])

	form.Options = make(Map)
	if options, ok := data["options"].(map[string]interface{}); ok {
		for key, value := range options {
			form.Options[key] = convert.String(value)
		}
	}

	if children, ok := data["children"].([]interface{}); ok {
		form.Children = make([]Form, len(children))
		for index, childInterface := range children {
			if childData, ok := childInterface.(map[string]interface{}); ok {
				var childForm Form
				childForm.UnmarshalMap(childData)
				form.Children[index] = childForm
			} else {
				return derp.NewInternalError("form.UnmarshalMap", "Error parsing child form information.", childInterface)
			}
		}
	}

	return nil
}

// 	Autocomplete string `json:"autocomplete"` // https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/autocomplete

// HTML returns a populated HTML string for the provided form, schema, and value
func (form Form) HTML(library *Library, schema *schema.Schema, value interface{}) (string, error) {

	b := html.New()

	if err := form.Write(library, schema, value, b); err != nil {
		return "", derp.Wrap(err, "form.HTML", "Error rendering element", form)
	}

	return b.String(), nil
}

// Write generates an HTML string for the fully populated form into the provided string builder
func (form Form) Write(library *Library, schema *schema.Schema, value interface{}, b *html.Builder) error {

	// Try to locate the Renderer in the library
	renderer, err := library.Renderer(form.Kind)

	if err != nil {
		return derp.Wrap(err, "form.Write", "Renderer Not Defined", form)
	}

	// try to render the form into the
	if err := renderer(form, schema, value, b); err != nil {
		return derp.Wrap(err, "form.Write", "Error rendering element", form)
	}

	return nil
}

// AllPaths returns pointers to all of the valid paths in this form
func (form Form) AllPaths() []Form {

	var result []Form

	// If THIS element has a Path, then add it to the result
	if form.Path != "" {
		result = []Form{form}
	} else {
		result = []Form{}
	}

	// Scan all chiild elements for THEIR paths, and add them to the result
	for _, child := range form.Children {
		result = append(result, child.AllPaths()...)
	}

	// Success
	return result
}

/*********************************
 * BUILDER FUNCTIONS
 *********************************/

// SetPath sets the path value for this form item
// This DOES NOT implement the path.Setter interface
func (form *Form) SetPath(path string) *Form {
	form.Path = path
	return form
}

// SetID sets the ID value for this form item.
func (form *Form) SetID(id string) *Form {
	form.ID = id
	return form
}

// SetLabel sets the Label value for this form item
func (form *Form) SetLabel(label string) *Form {
	form.Label = label
	return form
}

// SetDescription sets the description value for this form item
func (form *Form) SetDescription(description string) *Form {
	form.Description = description
	return form
}

// SetCSSClass sets the CSSClass value for this form item
func (form *Form) SetCSSClass(cssClass string) *Form {
	form.CSSClass = cssClass
	return form
}

// SetOption sets an individual option for this form item
func (form *Form) SetOption(key string, value string) *Form {
	form.Options[key] = value
	return form
}

// AddChild adds a new child of the designated type
// to this form element.  It returns a reference to the
// newly created child.
func (form *Form) AddChild(kind string) *Form {

	child := NewForm(kind)

	form.Children = append(form.Children, child)

	return &(form.Children[len(form.Children)-1])
}
