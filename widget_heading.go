package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("heading", WidgetHeading{})
}

type WidgetHeading struct{}

func (WidgetHeading) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := element.GetString(value, s)
	b.H2().InnerText(valueString).Close()
	return nil
}

func (WidgetHeading) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := element.GetString(value, s)
	b.H2().InnerText(valueString).Close()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (WidgetHeading) ShowLabels() bool {
	return false
}
