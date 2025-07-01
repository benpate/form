package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

type LayoutGroup struct{}

func (LayoutGroup) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(f, e, provider, value, b, "group", false)
}

func (LayoutGroup) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(f, e, provider, value, b, "group", true)
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (LayoutGroup) ShowLabels() bool {
	return false
}

func (widget LayoutGroup) Encoding(element *form.Element) string {
	return collectEncoding(element.Children)
}
