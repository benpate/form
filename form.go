package form

import (
	"github.com/benpate/derp"
	"github.com/benpate/form/html"
	"github.com/benpate/schema"
)

// Form defines a single form element, or a nested form layout.  It can be serialized to and from a database.
type Form struct {
	Kind        string                 `json:"kind"`        // The kind of form element
	ID          string                 `json:"id"`          // DOM ID to use for this element.
	Label       string                 `json:"label"`       // Short label to be displayed on the form element
	Description string                 `json:"description"` // Longer description text to be displayed on the form element
	CSSClass    string                 `json:"cssClass"`    // CSS Class override to apply to this widget.  This should be used sparingly
	Path        string                 `json:"path"`        // Path to the data value displayed in for this form element
	Options     map[string]interface{} `json:"options"`     // Kind-specific modifiers for the form element.
	Children    []Form                 `json:"children"`    // Array of sub-form elements that may be displayed depending on the kind.
}

// 	Autocomplete string `json:"autocomplete"` // https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/autocomplete

// HTML returns a populated HTML string for the provided form, schema, and value
func (form Form) HTML(library Library, schema *schema.Schema, value interface{}) (string, error) {

	builder := html.New()

	if err := form.Write(library, schema, value, builder); err != nil {
		return "", derp.Wrap(err, "form.HTML", "Error rendering element", form)
	}

	return builder.String(), nil
}

// Write generates an HTML string for the fully populated form into the provided string builder
func (form Form) Write(library Library, schema *schema.Schema, value interface{}, builder *html.Builder) error {

	// Try to locate the Renderer in the library
	renderer, err := library.Renderer(form.Kind)

	if err != nil {
		return derp.Wrap(err, "form.Write", "Renderer Not Defined", form)
	}

	// try to render the form into the
	if err := renderer(form, schema, value, builder); err != nil {
		return derp.Wrap(err, "form.Write", "Error rendering element", form)
	}

	return nil
}

// AllPaths returns pointers to all of the valid paths in this form
func (form Form) AllPaths() []*Form {

	var result []*Form

	// If THIS element has a Path, then add it to the result
	if form.Path != "" {
		result = []*Form{&form}
	} else {
		result = []*Form{}
	}

	// Scan all chiild elements for THEIR paths, and add them to the result
	for _, child := range form.Children {
		result = append(result, child.AllPaths()...)
	}

	// Success
	return result
}
