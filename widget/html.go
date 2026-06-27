package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

// HTML is a widget that emits the element's Description as raw, unescaped HTML.
type HTML struct{}

// View generates the read-only HTML for this element from its Description.
func (widget HTML) View(_ *form.Form, e *form.Element, _ form.LookupProvider, _ any, b *html.Builder) error {
	b.Div().InnerHTML(e.Description).Close()
	return nil
}

// Edit generates the HTML for this element, which is identical to its read-only view.
func (widget HTML) Edit(_ *form.Form, e *form.Element, _ form.LookupProvider, _ any, b *html.Builder) error {
	b.Div().InnerHTML(e.Description).Close()
	return nil
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For HTML widgets, labels are not shown, so this always returns FALSE.
func (widget HTML) ShowLabels() bool {
	return false
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget HTML) ShowDescriptions() string {
	return "NONE"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For HTML widgets, there is no special encoding,
// so this always returns an empty string.
func (widget HTML) Encoding(_ *form.Element) string {
	return ""
}
