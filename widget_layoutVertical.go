package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("layout-vertical", WidgetLayoutVertical{})
}

type WidgetLayoutVertical struct{}

func (WidgetLayoutVertical) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	return drawLayout(element, s, lookupProvider, value, b, "vertical", false)
}

func (WidgetLayoutVertical) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetLayoutVertical{}.View(element, s, lookupProvider, value, b)
	}

	return drawLayout(element, s, lookupProvider, value, b, "vertical", true)
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (WidgetLayoutVertical) ShowLabels() bool {
	return false
}
