package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

type Colorpicker struct{}

func (widget Colorpicker) View(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := element.GetString(value, s)

	// TODO: LOW: Apply formatting options?
	b.Div().Class("layout-value", element.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

func (widget Colorpicker) Edit(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return Colorpicker{}.View(element, s, nil, value, b)
	}

	// find the path and schema to use
	valueString := element.GetString(value, s)

	// Start building a new tag
	b.Div().
		Data("path", element.Path).
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
