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

	if valueBool := convert.Bool(toggleValue(f, e, value)); valueBool {
		b.Div().Class("layout-value").InnerText(e.Options.GetString("true-text")).Close()
	} else {
		b.Div().Class("layout-value").InnerText(e.Options.GetString("false-text")).Close()
	}

	return nil
}

// Edit generates the HTML for editing a toggle widget.
func (widget Toggle) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := toggleValue(f, e, value)
	id := e.ID
	if id == "" {
		id = "toggle-" + strings.ReplaceAll(e.Path, ".", "-") + "-" + valueString
	}

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

// toggleValue reads the toggle's value from the object, falling back to the schema's declared
// default when the object carries no value for this path at all.
//
// RULE: An absent property and a stored FALSE are NOT the same thing, but they collapse into
// the same rendering unless the default is applied here.  A map-backed object simply has no
// key yet, so the value reads back as an empty string -- and because a toggle always posts
// either "true" or "false", the very first save of the surrounding form would write that
// phantom FALSE into storage.  A flag whose ON state is meant to be the default therefore has
// to render ON before it has ever been saved, which is what this restores.
func toggleValue(f *form.Form, e *form.Element, value any) string {

	if result := e.GetString(value, &f.Schema); result != "" {
		return result
	}

	if element := e.GetSchema(&f.Schema); element != nil {
		return convert.String(element.DefaultValue())
	}

	return ""
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For Toggle widgets, labels are shown, so this always returns TRUE.
func (widget Toggle) ShowLabels() bool {
	return true
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget Toggle) ShowDescriptions() string {
	return "TOP"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For Toggle widgets, there is no special encoding,
// so this always returns an empty string.
func (widget Toggle) Encoding(_ *form.Element) string {
	return ""
}
