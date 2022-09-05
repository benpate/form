package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("layout-horizontal", WidgetLayoutHorizontal{})
}

type WidgetLayoutHorizontal struct{}

func (WidgetLayoutHorizontal) View(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, schema, lookupProvider, value, b, "horizontal", false)
}

func (WidgetLayoutHorizontal) Edit(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, schema, lookupProvider, value, b, "horizontal", true)
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (WidgetLayoutHorizontal) ShowLabels() bool {
	return false
}
