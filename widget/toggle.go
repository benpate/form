package widget

import (
	"strings"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
)

// Toggle renders a custom toggle widget
type Toggle struct{}

func (widget Toggle) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	valueString := e.GetString(value, &f.Schema)

	if valueBool := convert.Bool(valueString); valueBool {
		b.Div().Class("layout-value").InnerText(e.Options.GetString("true-text")).Close()
	} else {
		b.Div().Class("layout-value").InnerText(e.Options.GetString("false-text")).Close()
	}

	return nil
}

func (widget Toggle) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)
	id := "toggle-" + strings.ReplaceAll(e.Path, ".", "-") + "-" + valueString
	script := "install toggle " + e.Options.GetString("script")

	// Start building a new tag
	tag := b.Span().ID(id).Script(script).Name(e.Path)

	if convert.Bool(valueString) {
		tag.Value("true")
	}

	tag.Attr("text", e.Options.GetString("text"))
	tag.Attr("true-text", e.Options.GetString("true-text"))
	tag.Attr("false-text", e.Options.GetString("false-text"))

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
