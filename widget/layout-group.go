package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

type LayoutGroup struct{}

func (LayoutGroup) View(element *form.Element, schema *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, schema, provider, value, b, "group", false)
}

func (LayoutGroup) Edit(element *form.Element, schema *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, schema, provider, value, b, "group", true)
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
