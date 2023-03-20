package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("label", WidgetLabel{})
}

type WidgetLabel struct{}

func (WidgetLabel) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := element.GetString(value, s)
	b.Div().InnerText(valueString).Close()
	return nil
}

func (WidgetLabel) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := element.GetString(value, s)
	b.Div().InnerText(valueString).Close()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (WidgetLabel) ShowLabels() bool {
	return false
}
