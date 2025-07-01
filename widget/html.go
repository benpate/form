package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

type HTML struct{}

func (widget HTML) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	b.Div().InnerHTML(e.Description).Close()
	return nil
}

func (widget HTML) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	b.Div().InnerHTML(e.Description).Close()
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
