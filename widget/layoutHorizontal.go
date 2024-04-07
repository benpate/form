package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("layout-horizontal", WidgetLayoutHorizontal{})
}

type WidgetLayoutHorizontal struct{}

func (widget WidgetLayoutHorizontal) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, s, lookupProvider, value, b, "horizontal", false)
}

func (widget WidgetLayoutHorizontal) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetLayoutHorizontal{}.View(element, s, lookupProvider, value, b)
	}

	return drawLayout(element, s, lookupProvider, value, b, "horizontal", true)
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget WidgetLayoutHorizontal) ShowLabels() bool {
	return false
}

func (widget WidgetLayoutHorizontal) Encoding(element *Element) string {
	return collectEncoding(element.Children)
}
