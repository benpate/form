package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

type LayoutVertical struct{}

func (widget LayoutVertical) View(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, s, provider, value, b, "vertical", false)
}

func (widget LayoutVertical) Edit(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return drawLayout(element, s, provider, value, b, "vertical", false)
	}

	return drawLayout(element, s, provider, value, b, "vertical", true)
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
