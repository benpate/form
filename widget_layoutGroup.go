package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("layout-group", WidgetLayoutGroup{})
}

type WidgetLayoutGroup struct{}

func (WidgetLayoutGroup) View(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, schema, lookupProvider, value, b, "group", false)
}

func (WidgetLayoutGroup) Edit(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, schema, lookupProvider, value, b, "group", true)
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (WidgetLayoutGroup) ShowLabels() bool {
	return false
}
