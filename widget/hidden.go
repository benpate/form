package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
)

// Hidden is a widget that creates a hidden input field.
type Hidden struct{}

// View generates the HTML for viewing a hidden input field (which is nothing).
func (widget Hidden) View(_ *form.Form, _ *form.Element, _ form.LookupProvider, _ any, _ *html.Builder) error {
	return nil
}

// Edit generates the HTML for editing a hidden input field.
func (widget Hidden) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	var elementValue string

	if optionValue, ok := e.Options["value"]; ok {
		elementValue = convert.String(optionValue)
	} else {
		elementValue = e.GetString(value, &f.Schema)
	}

	// Start building a new tag
	b.Input("hidden", e.Path).
		ID(e.Path).
		Value(elementValue).
		Close()

	return nil
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For Hidden widgets, nothing is visible, so this always returns FALSE.
func (widget Hidden) ShowLabels() bool {
	return false
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget Hidden) ShowDescriptions() string {
	return "NONE"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For Hidden widgets, there is no special encoding,
// so this always returns an empty string.
func (widget Hidden) Encoding(_ *form.Element) string {
	return ""
}
