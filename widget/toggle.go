package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("toggle", WidgetToggle{})
}

// WidgetToggle renders a custom toggle widget
type WidgetToggle struct{}

func (widget WidgetToggle) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	valueString := element.GetString(value, s)
	valueBool := convert.Bool(valueString)

	if valueBool {
		b.Div().Class("layout-value").InnerText(element.Options.GetString("true-text")).Close()
	} else {
		b.Div().Class("layout-value").InnerText(element.Options.GetString("false-text")).Close()
	}

	return nil
}

func (widget WidgetToggle) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetToggle{}.View(element, s, lookupProvider, value, b)
	}

	// find the path and schema to use
	valueString := element.GetString(value, s)

	// Start building a new tag
	tag := b.Span().Script("install toggle").Name(element.Path)

	if convert.Bool(valueString) {
		tag.Value("true")
	}

	tag.Attr("true-text", element.Options.GetString("true-text"))
	tag.Attr("false-text", element.Options.GetString("false-text"))

	b.CloseAll()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget WidgetToggle) ShowLabels() bool {
	return true
}

func (widget WidgetToggle) Encoding(_ *Element) string {
	return ""
}
