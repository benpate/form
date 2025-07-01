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

	if e.ReadOnly {
		return drawLayout(f, e, provider, value, b, "vertical", false)
	}

	return drawLayout(f, e, provider, value, b, "vertical", true)
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget LayoutVertical) ShowLabels() bool {
	return false
}

func (widget LayoutVertical) Encoding(element *form.Element) string {
	return collectEncoding(element.Children)
}
