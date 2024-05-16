package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

type Label struct{}

func (widget Label) View(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {
	if element.Label != "" {
		b.Div().InnerText(element.Label).Close()
	}

	if element.Description != "" {
		b.Div().InnerHTML(element.Description).Close()
	}

	return nil
}

func (widget Label) Edit(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {
	if element.Label != "" {
		b.Div().InnerText(element.Label).Close()
	}

	if element.Description != "" {
		b.Div().InnerHTML(element.Description).Close()
	}

	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget Label) ShowLabels() bool {
	return false
}

func (widget Label) Encoding(_ *form.Element) string {
	return ""
}
