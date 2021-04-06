package form

import (
	"github.com/benpate/derp"
	"github.com/benpate/html"
	"github.com/benpate/schema"
)

// Form defines a single form element, or a nested form layout.  It can be serialized to and from a database.
type Form struct {
	Path        string            `json:"path"`        // Path to the data value displayed in for this form element
	Kind        string            `json:"kind"`        // The kind of form element
	Widget      string            `json:"widget"`      // Optional: used to specify the kind of widget to use instead of the default
	ID          string            `json:"id"`          // DOM ID to use for this element.
	Label       string            `json:"label"`       // Short label to be displayed on the form element
	Description string            `json:"description"` // Longer description text to be displayed on the form element
	Options     string            `json:"options"`     // URL of the OptionProvider for this widget.
	CSSClass    string            `json:"cssClass"`    // CSS Class override to apply to this widget.  This should be used sparingly
	Children    []Form            `json:"children"`    // Array of sub-form elements that may be displayed depending on the kind.
	Rules       map[string]string `json:"rules"`       // Visibility rules (in hyperscript) to apply to UI.
}

// 	Autocomplete string `json:"autocomplete"` // https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/autocomplete

// HTML returns a populated HTML string for the provided form, schema, and value
func (form Form) HTML(library Library, schema *schema.Schema, value interface{}) (string, error) {

	b := html.New()

	if err := form.Write(library, schema, value, b); err != nil {
		return "", derp.Wrap(err, "form.HTML", "Error rendering element", form)
	}

	return b.String(), nil
}

// Write generates an HTML string for the fully populated form into the provided string builder
func (form Form) Write(library Library, schema *schema.Schema, value interface{}, b *html.Builder) error {

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
