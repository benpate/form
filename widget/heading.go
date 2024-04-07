package widget

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("heading", Heading{})
}

type Heading struct{}

func (widget Heading) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := element.GetString(value, s)
	b.H2().InnerText(valueString).Close()
	return nil
}

func (widget Heading) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := element.GetString(value, s)
	b.H2().InnerText(valueString).Close()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget Heading) ShowLabels() bool {
	return false
}

func (widget Heading) Encoding(_ *Element) string {
	return ""
}
