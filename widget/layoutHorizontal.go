package widget

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("layout-horizontal", LayoutHorizontal{})
}

type LayoutHorizontal struct{}

func (widget LayoutHorizontal) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, s, lookupProvider, value, b, "horizontal", false)
}

func (widget LayoutHorizontal) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return LayoutHorizontal{}.View(element, s, lookupProvider, value, b)
	}

	return drawLayout(element, s, lookupProvider, value, b, "horizontal", true)
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget LayoutHorizontal) ShowLabels() bool {
	return false
}

func (widget LayoutHorizontal) Encoding(element *Element) string {
	return collectEncoding(element.Children)
}
