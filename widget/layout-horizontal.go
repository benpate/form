package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

// LayoutHorizontal is a widget that arranges its child elements in a horizontal row.
type LayoutHorizontal struct{}

// View generates the read-only HTML for this horizontal layout.
func (widget LayoutHorizontal) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(f, e, provider, value, b, "horizontal", false)
}

// Edit generates the editable HTML for this horizontal layout.
func (widget LayoutHorizontal) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(f, e, provider, value, b, "horizontal", true)
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For LayoutHorizontal widgets, labels are not shown, so this always returns FALSE.
func (widget LayoutHorizontal) ShowLabels() bool {
	return false
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget LayoutHorizontal) ShowDescriptions() string {
	return "NONE"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget,
// which is the combined encoding of all child elements.
func (widget LayoutHorizontal) Encoding(element *form.Element) string {
	return collectEncoding(element.Children)
}
