package widget

import (
	"strings"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
)

// Toggle renders a custom toggle widget
type Toggle struct{}

// View generates the HTML for viewing a toggle widget's value.
func (widget Toggle) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	valueString := e.GetString(value, &f.Schema)

	if valueBool := convert.Bool(valueString); valueBool {
		b.Div().Class("layout-value").InnerText(e.Options.GetString("true-text")).Close()
	} else {
		b.Div().Class("layout-value").InnerText(e.Options.GetString("false-text")).Close()
	}

	return nil
}

// Edit generates the HTML for editing a toggle widget.
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

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For Toggle widgets, labels are shown, so this always returns TRUE.
func (widget Toggle) ShowLabels() bool {
	return true
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For Toggle widgets, there is no special encoding,
// so this always returns an empty string.
func (widget Toggle) Encoding(_ *form.Element) string {
	return ""
}
