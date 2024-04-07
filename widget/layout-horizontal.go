package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

type LayoutHorizontal struct{}

func (widget LayoutHorizontal) View(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, s, provider, value, b, "horizontal", false)
}

func (widget LayoutHorizontal) Edit(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return LayoutHorizontal{}.View(element, s, provider, value, b)
	}

	return drawLayout(element, s, provider, value, b, "horizontal", true)
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
