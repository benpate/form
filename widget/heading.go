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
	b.H2().InnerHTML(e.Label).Close()
	b.Div().InnerHTML(e.Description).Close()
	return nil
}

/*******************************************
 * Widget Metadata
 *******************************************/

func (widget Heading) ShowLabels() bool {
	return false
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget Heading) ShowDescriptions() string {
	return "NONE"
}

func (widget Heading) Encoding(_ *form.Element) string {
	return ""
}
