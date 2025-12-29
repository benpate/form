package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

// Colorpicker is a widget that creates a color picker input field.
type Colorpicker struct{}

// View generates the HTML for viewing a color picker value.
func (widget Colorpicker) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)

	// TODO: LOW: Apply formatting options?
	b.Div().Class("layout-value", e.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

// Edit generates the HTML for editing a color picker input field.
func (widget Colorpicker) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)

	// Start building a new tag
	b.Div().
		Data("label", e.Label).
		Data("description", e.Description).
		Data("path", e.Path).
		Data("value", valueString).
		Script("install colorpicker").
		Close()

	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For Colorpicker widgets, labels are shown, so this always returns TRUE.
func (widget Colorpicker) ShowLabels() bool {
	return true
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For Colorpicker widgets, there is no special encoding,
// so this always returns an empty string.
func (widget Colorpicker) Encoding(_ *form.Element) string {
	return ""
}
