package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

// Toggle renders a custom toggle widget
type Toggle struct{}

func (widget Toggle) View(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {

	valueString := element.GetString(value, s)
	valueBool := convert.Bool(valueString)

	if valueBool {
		b.Div().Class("layout-value").InnerText(element.Options.GetString("true-text")).Close()
	} else {
		b.Div().Class("layout-value").InnerText(element.Options.GetString("false-text")).Close()
	}

	return nil
}

func (widget Toggle) Edit(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return Toggle{}.View(element, s, nil, value, b)
	}

	// find the path and schema to use
	valueString := element.GetString(value, s)

	// Start building a new tag
	tag := b.Span().Script("install toggle").Name(element.Path)

	if convert.Bool(valueString) {
		tag.Value("true")
	}

	tag.Attr("text", element.Options.GetString("text"))
	tag.Attr("true-text", element.Options.GetString("true-text"))
	tag.Attr("false-text", element.Options.GetString("false-text"))

	b.CloseAll()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget Toggle) ShowLabels() bool {
	return true
}

func (widget Toggle) Encoding(_ *form.Element) string {
	return ""
}
