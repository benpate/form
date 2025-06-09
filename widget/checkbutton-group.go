package widget

import (
	"strings"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
)

// CheckButtonGroup renders a group of fancy checkboxes that look like buttons
type CheckButtonGroup struct{}

func (widget CheckButtonGroup) View(element *form.Element, schema *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the schema and value to use
	schemaElement := element.GetSchema(schema)
	valueSlice := element.GetSliceOfString(value, schema)
	lookupCodes := widget.getLookupCodes(element, schemaElement, provider)

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

func (widget CheckButtonGroup) Edit(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return CheckButtonGroup{}.View(element, s, provider, value, b)
	}

	// find the path and schema to use
	schemaElement := element.GetSchema(s)
	valueSlice := element.GetSliceOfString(value, s)
	lookupCodes := widget.getLookupCodes(element, schemaElement, provider)

	// Start building a new tag
	for _, lookupCode := range lookupCodes {
		id := "checkbutton-" + strings.ReplaceAll(element.Path, ".", "-") + "-" + lookupCode.Value

		b.Label(id).ID("label-" + id).Class("checkbutton")

		if lookupCode.Icon != "" {
			b.I().Class("margin-horizontal", "bi", "bi-"+lookupCode.Icon).Style("font-size:32px;").Close()
		}

		b.Div().Class("flex-column")
		b.Div().Class("bold").InnerText(lookupCode.Label).Close()
		b.Div().Class("text-sm", "text-gray").InnerText(lookupCode.Description).Close()

		toggleButton := b.Input("checkbox", element.Path).
			ID(id).
			Class(element.Options.GetString("class")).
			Value(lookupCode.Value).
			Aria("label", lookupCode.Label).
			Aria("description", lookupCode.Description).
			Script(element.Options.GetString("script")).
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

// getLookupCodes returns a list of LookupCodes for this element
func (widget CheckButtonGroup) getLookupCodes(element *form.Element, schemaElement schema.Element, provider form.LookupProvider) []form.LookupCode {

	lookupCodes, _ := form.GetLookupCodes(element, schemaElement, provider)

	if len(lookupCodes) == 0 {
		lookupCodes = []form.LookupCode{
			{Value: "true", Label: element.Label},
		}
	}

	return lookupCodes
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget CheckButtonGroup) ShowLabels() bool {
	return true
}

func (widget CheckButtonGroup) Encoding(_ *form.Element) string {
	return ""
}
