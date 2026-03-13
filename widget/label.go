package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

type Label struct{}

func (widget Label) View(_ *form.Form, e *form.Element, _ form.LookupProvider, _ any, b *html.Builder) error {
	if e.Label != "" {
		b.Div().InnerText(e.Label).Close()
	}

	if e.Description != "" {
		b.Div().InnerHTML(e.Description).Close()
	}

	return nil
}

func (widget Label) Edit(_ *form.Form, e *form.Element, _ form.LookupProvider, _ any, b *html.Builder) error {
	if e.Label != "" {
		b.Div().InnerText(e.Label).Close()
	}

	if e.Description != "" {
		b.Div().InnerHTML(e.Description).Close()
	}

	return nil
}

/***********************************
 * Widget Metadata
 ***********************************/

func (widget Label) ShowLabels() bool {
	return false
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget Label) ShowDescriptions() string {
	return "NONE"
}

func (widget Label) Encoding(_ *form.Element) string {
	return ""
}
