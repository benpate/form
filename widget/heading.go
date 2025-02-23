package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

type Heading struct{}

func (widget Heading) View(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {
	b.H2().InnerText(element.Label).Close()
	b.Div().InnerHTML(element.Description).Close()
	return nil
}

func (widget Heading) Edit(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {
	b.H2().InnerText(element.Label).Close()
	b.Div().InnerHTML(element.Description).Close()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget Heading) ShowLabels() bool {
	return false
}

func (widget Heading) Encoding(_ *form.Element) string {
	return ""
}
