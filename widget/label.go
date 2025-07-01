package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

type Label struct{}

func (widget Label) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	if e.Label != "" {
		b.Div().InnerText(e.Label).Close()
	}

	if e.Description != "" {
		b.Div().InnerHTML(e.Description).Close()
	}

	return nil
}

func (widget Label) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	if e.Label != "" {
		b.Div().InnerText(e.Label).Close()
	}

	if e.Description != "" {
		b.Div().InnerHTML(e.Description).Close()
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
