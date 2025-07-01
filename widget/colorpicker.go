package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

type Colorpicker struct{}

func (widget Colorpicker) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)

	// TODO: LOW: Apply formatting options?
	b.Div().Class("layout-value", e.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

func (widget Colorpicker) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	if e.ReadOnly {
		return Colorpicker{}.View(f, e, provider, value, b)
	}

	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)

	// Start building a new tag
	b.Div().
		Data("label", e.Label).
		Data("description", e.Description).
		Data("path", e.Path).
		Data("value", valueString).
		Script("install colorpicker").
		Close()

	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget Colorpicker) ShowLabels() bool {
	return true
}

func (widget Colorpicker) Encoding(_ *form.Element) string {
	return ""
}
