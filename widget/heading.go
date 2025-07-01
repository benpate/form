package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

type Heading struct{}

func (widget Heading) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	b.H2().InnerText(e.Label).Close()
	b.Div().InnerHTML(e.Description).Close()
	return nil
}

func (widget Heading) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	b.H2().InnerText(e.Label).Close()
	b.Div().InnerHTML(e.Description).Close()
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
