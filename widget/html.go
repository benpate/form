package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

type HTML struct{}

func (widget HTML) View(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {
	b.Div().InnerHTML(element.Description).Close()
	return nil
}

func (widget HTML) Edit(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {
	b.Div().InnerHTML(element.Description).Close()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget HTML) ShowLabels() bool {
	return false
}

func (widget HTML) Encoding(_ *form.Element) string {
	return ""
}
