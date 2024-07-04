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

	elementID := element.ID

	if elementID == "" {
		elementID = "colorpicker-" + element.Path
	}

	// Start building a new tag
	tag := b.Input("", element.Path).
		ID(elementID)

	if focus, ok := element.Options.GetBoolOK("focus"); ok && focus {
		tag.Attr("autofocus", "true")
	}

	if placeholder := element.Options.GetString("placeholder"); placeholder != "" {
		tag.Attr("placeholder", placeholder)
	}

	tag.TabIndex("0")
	tag.Type("color")
	tag.Value(valueString)
	tag.Close()

	b.CloseAll()
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
