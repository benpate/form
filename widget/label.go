package widget

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("label", Label{})
}

type Label struct{}

func (widget Label) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := element.GetString(value, s)
	b.Div().InnerText(valueString).Close()
	return nil
}

func (widget Label) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := element.GetString(value, s)
	b.Div().InnerText(valueString).Close()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget Label) ShowLabels() bool {
	return false
}

func (widget Label) Encoding(_ *Element) string {
	return ""
}
