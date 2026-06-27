package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

// LayoutVertical is a widget that arranges its child elements in a vertical stack.
type LayoutVertical struct{}

// View generates the read-only HTML for this vertical layout.
func (widget LayoutVertical) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(f, e, provider, value, b, "vertical", false)
}

// Edit generates the editable HTML for this vertical layout.
func (widget LayoutVertical) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(f, e, provider, value, b, "vertical", true)
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For LayoutVertical widgets, labels are not shown, so this always returns FALSE.
func (widget LayoutVertical) ShowLabels() bool {
	return false
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget LayoutVertical) ShowDescriptions() string {
	return "NONE"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget,
// which is the combined encoding of all child elements.
func (widget LayoutVertical) Encoding(element *form.Element) string {
	return collectEncoding(element.Children)
}
