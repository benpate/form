package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

// Heading is a widget that displays a section heading and description, with no input.
type Heading struct{}

// View generates the read-only HTML for this heading.
func (widget Heading) View(_ *form.Form, e *form.Element, _ form.LookupProvider, _ any, b *html.Builder) error {
	b.H2().InnerText(e.Label).Close()
	b.Div().InnerHTML(e.Description).Close()
	return nil
}

// Edit generates the HTML for this heading, which is identical to its read-only view.
func (widget Heading) Edit(_ *form.Form, e *form.Element, _ form.LookupProvider, _ any, b *html.Builder) error {
	b.H2().InnerHTML(e.Label).Close()
	b.Div().InnerHTML(e.Description).Close()
	return nil
}

/*******************************************
 * Widget Metadata
 *******************************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For Heading widgets, the heading is the content itself, so this always returns FALSE.
func (widget Heading) ShowLabels() bool {
	return false
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget Heading) ShowDescriptions() string {
	return "NONE"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For Heading widgets, there is no special encoding,
// so this always returns an empty string.
func (widget Heading) Encoding(_ *form.Element) string {
	return ""
}
