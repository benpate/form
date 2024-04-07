package widget

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("layout-vertical", LayoutVertical{})
}

type LayoutVertical struct{}

func (widget LayoutVertical) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, s, lookupProvider, value, b, "vertical", false)
}

func (widget LayoutVertical) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return drawLayout(element, s, lookupProvider, value, b, "vertical", false)
	}

	return drawLayout(element, s, lookupProvider, value, b, "vertical", true)
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget LayoutVertical) ShowLabels() bool {
	return false
}

func (widget LayoutVertical) Encoding(element *Element) string {
	return collectEncoding(element.Children)
}
