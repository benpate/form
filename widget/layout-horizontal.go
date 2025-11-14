package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

type LayoutHorizontal struct{}

func (widget LayoutHorizontal) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(f, e, provider, value, b, "horizontal", false)
}

func (widget LayoutHorizontal) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(f, e, provider, value, b, "horizontal", true)
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget LayoutHorizontal) ShowLabels() bool {
	return false
}

func (widget LayoutHorizontal) Encoding(element *form.Element) string {
	return collectEncoding(element.Children)
}
