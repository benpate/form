package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

// LayoutGroup is a widget that arranges its child elements as a labeled group.
type LayoutGroup struct{}

// View generates the read-only HTML for this group layout.
func (LayoutGroup) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(f, e, provider, value, b, "group", false)
}

// Edit generates the editable HTML for this group layout.
func (LayoutGroup) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(f, e, provider, value, b, "group", true)
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For LayoutGroup widgets, labels are not shown, so this always returns FALSE.
func (LayoutGroup) ShowLabels() bool {
	return false
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget LayoutGroup) ShowDescriptions() string {
	return "NONE"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget,
// which is the combined encoding of all child elements.
func (widget LayoutGroup) Encoding(element *form.Element) string {
	return collectEncoding(element.Children)
}
