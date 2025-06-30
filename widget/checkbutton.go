package widget

import (
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/compare"
	"github.com/benpate/rosetta/schema"
)

// CheckButton renders a fancy checkbox widget that looks like a button
type CheckButton struct{}

func (widget CheckButton) View(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {
	return nil
}

func (widget CheckButton) Edit(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return CheckButton{}.View(element, s, nil, value, b)
	}

	// Collect values to use in the Widget
	selectedValues, err := s.Get(value, element.Path)

	if err != nil {
		derp.Report(derp.Wrap(err, "form.checkbutton.Edit", "Error getting value for CheckButton", element.Path, value))
	}

	elementValue := element.Options.GetString("value")
	id := "checkbutton-" + strings.ReplaceAll(element.Path, ".", "-") + "-" + elementValue

	// Build the widget HTML
	b.Label(id).ID("label-" + id).Class("checkbutton")

	if icon := element.Options.GetString("icon"); icon != "" {
		b.I().Class("margin-horizontal", "bi", "bi-"+icon).Style("font-size:32px;").Close()
	}

	b.Div().Class("flex-column")
	b.Div().Class("bold").InnerText(element.Label).Close()
	b.Div().Class("text-sm", "text-gray").InnerText(element.Description).Close()

	checkbox := b.Input("checkbox", element.Path).
		ID(id).
		Value(elementValue).
		Class(element.Options.GetString("class")).
		Script(element.Options.GetString("script")).
		Aria("label", element.Label).
		Aria("description", element.Description).
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
