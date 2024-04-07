package widget

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("layout-group", LayoutGroup{})
}

type LayoutGroup struct{}

func (LayoutGroup) View(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, schema, lookupProvider, value, b, "group", false)
}

func (LayoutGroup) Edit(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, schema, lookupProvider, value, b, "group", true)
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (LayoutGroup) ShowLabels() bool {
	return false
}

func (widget LayoutGroup) Encoding(element *Element) string {
	return collectEncoding(element.Children)
}
