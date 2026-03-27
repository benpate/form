package widget

import (
	"strings"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/slice"
)

// CheckButtonGroup renders a group of fancy checkboxes that look like buttons
type CheckButtonGroup struct{}

func (widget CheckButtonGroup) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the schema and value to use
	schemaElement := e.GetSchema(&f.Schema)
	valueSlice := e.GetSliceOfString(value, &f.Schema)
	lookupCodes, _ := form.GetLookupCodes(e, schemaElement, provider)

	first := true

	b.Div().Class("layout-value")
	for _, lookupCode := range lookupCodes {

		if slice.Contains(valueSlice, lookupCode.Value) {

			if first {
				first = false
			} else {
				b.WriteString(", ")
			}

			b.WriteString(lookupCode.Label)
		}
	}
	b.Close()

	return nil
}

func (widget CheckButtonGroup) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)
	valueSlice := e.GetSliceOfString(value, &f.Schema)
	lookupCodes, _ := form.GetLookupCodes(e, schemaElement, provider)

	// Start building a new tag
	for _, lookupCode := range lookupCodes {
		id := "checkbutton-" + strings.ReplaceAll(e.Path, ".", "-") + "-" + lookupCode.Value

		b.Label(id).ID("label-" + id).Class("checkbutton")

		if lookupCode.Icon != "" {
			b.I().Class("margin-horizontal", "bi", "bi-"+lookupCode.Icon).Style("font-size:32px;").Close()
		}

		b.Div().Class("flex-column")
		b.Div().Class("bold").InnerText(lookupCode.Label).Close()
		b.Div().Class("text-sm", "text-gray").InnerText(lookupCode.Description).Close()

		toggleButton := b.Input("checkbox", e.Path)

		toggleButton.
			ID(id).
			Class(e.Options.GetString("class")).
			Value(lookupCode.Value).
			Aria("label", lookupCode.Label).
			Aria("description", lookupCode.Description).
			Script(e.Options.GetString("script")).
			TabIndex("0")

		if slice.Contains(valueSlice, lookupCode.Value) {
			toggleButton.Attr("checked", "true")
		}

		b.Close()
		b.Close()
		b.Close()
	}

	return nil
}

/***********************************
 * Widget Metadata
 ***********************************/

func (widget CheckButtonGroup) ShowLabels() bool {

	return true
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget CheckButtonGroup) ShowDescriptions() string {
	return "TOP"
}

func (widget CheckButtonGroup) Encoding(_ *form.Element) string {
	return ""
}
