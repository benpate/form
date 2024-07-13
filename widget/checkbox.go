package widget

import (
	"strings"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
)

type Checkbox struct{}

func (widget Checkbox) View(element *form.Element, schema *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

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

func (widget Checkbox) Edit(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return Checkbox{}.View(element, s, provider, value, b)
	}

	// find the path and schema to use
	schemaElement := element.GetSchema(s)
	valueSlice := element.GetSliceOfString(value, s)
	lookupCodes := widget.getLookupCodes(element, schemaElement, provider)

	// Start building a new tag
	for _, lookupCode := range lookupCodes {
		id := "checkbox-" + strings.ReplaceAll(element.Path, ".", "-") + "-" + lookupCode.Value

		label := b.Label(id).ID("label-" + id)

		checkbox := b.Input("checkbox", element.Path).
			ID(id).
			Value(lookupCode.Value).
			Aria("label", lookupCode.Label).
			Aria("description", lookupCode.Description).
			TabIndex("0")

		if slice.Contains(valueSlice, lookupCode.Value) {
			checkbox.Attr("checked", "true")
		}
		checkbox.Close()

		b.Span().Aria("hidden", "true").InnerText(lookupCode.Label).Close()
		label.Close()
	}

	return nil
}

// getLookupCodes returns a list of LookupCodes for this element
func (widget Checkbox) getLookupCodes(element *form.Element, schemaElement schema.Element, provider form.LookupProvider) []form.LookupCode {

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

func (widget Checkbox) ShowLabels() bool {
	return false
}

func (widget Checkbox) Encoding(_ *form.Element) string {
	return ""
}
