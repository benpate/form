package widget

import (
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/compare"
)

// CheckButton renders a fancy checkbox widget that looks like a button
type CheckButton struct{}

func (widget CheckButton) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	return nil
}

func (widget CheckButton) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	// Collect values to use in the Widget
	selectedValues, err := f.Schema.Get(value, e.Path)

	if err != nil {
		derp.Report(derp.Wrap(err, "form.checkbutton.Edit", "Error getting value for CheckButton", e.Path, value))
	}

	elementValue := e.Options.GetString("value")
	id := "checkbutton-" + strings.ReplaceAll(e.Path, ".", "-") + "-" + elementValue

	// Build the widget HTML
	b.Label(id).ID("label-" + id).Class("checkbutton")

	if icon := e.Options.GetString("icon"); icon != "" {
		b.I().Class("margin-horizontal", "bi", "bi-"+icon).Style("font-size:32px;").Close()
	}

	b.Div().Class("flex-column")
	b.Div().Class("bold").InnerText(e.Label).Close()
	b.Div().Class("text-sm", "text-gray").InnerText(e.Description).Close()

	checkbox := b.Input("checkbox", e.Path).
		ID(id).
		Value(elementValue).
		Class(e.Options.GetString("class")).
		Script(e.Options.GetString("script")).
		Aria("label", e.Label).
		Aria("description", e.Description).
		TabIndex("0")

	if compare.Contains(selectedValues, elementValue) {
		checkbox.Attr("checked", "true")
	}

	b.CloseAll()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget CheckButton) ShowLabels() bool {
	return false
}

func (widget CheckButton) Encoding(_ *form.Element) string {
	return ""
}
