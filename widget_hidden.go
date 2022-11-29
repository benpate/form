package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("hidden", WidgetHidden{})
}

type WidgetHidden struct{}

func (WidgetHidden) View(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	return nil
}

// WidgetHidden registers a text <input> widget into the library
func (WidgetHidden) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetHidden{}.View(element, s, lookupProvider, value, b)
	}

	// find the path and schema to use
	elementValue := element.GetString(value, s)

	// Start building a new tag
	b.Input("hidden", element.Path).
		ID(element.Path).
		Value(elementValue).
		Close()

	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (WidgetHidden) ShowLabels() bool {
	return false
}
