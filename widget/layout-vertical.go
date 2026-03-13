package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

type LayoutVertical struct{}

func (widget LayoutVertical) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(f, e, provider, value, b, "vertical", false)
}

func (widget LayoutVertical) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(f, e, provider, value, b, "vertical", true)
}

/***********************************
 * Widget Metadata
 ***********************************/

func (widget LayoutVertical) ShowLabels() bool {
	return false
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget LayoutVertical) ShowDescriptions() string {
	return "NONE"
}

func (widget LayoutVertical) Encoding(element *form.Element) string {
	return collectEncoding(element.Children)
}
