package form

import (
	"strings"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
)

func init() {
	Register("checkbox", WidgetCheckbox{})
}

type WidgetCheckbox struct{}

func (widget WidgetCheckbox) View(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the schema and value to use
	schemaElement := element.getElement(schema)
	valueSlice := element.GetSliceOfString(value, schema)
	lookupCodes := widget.getLookupCodes(element, schemaElement, lookupProvider)

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

func (widget WidgetCheckbox) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetCheckbox{}.View(element, s, lookupProvider, value, b)
	}

	// find the path and schema to use
	schemaElement := element.getElement(s)
	valueSlice := element.GetSliceOfString(value, s)
	lookupCodes := widget.getLookupCodes(element, schemaElement, lookupProvider)

	// Start building a new tag
	for _, lookupCode := range lookupCodes {
		id := "checkbox-" + strings.ReplaceAll(element.Path, ".", "-") + "-" + lookupCode.Value
		b.Label(id)

		checkbox := b.Input("checkbox", element.Path).
			ID(id).
			Value(lookupCode.Value)

		if slice.Contains(valueSlice, lookupCode.Value) {
			checkbox.Attr("checked", "true")
		}

		checkbox.InnerHTML(lookupCode.Label).Close()
		b.CloseAll()
	}

	return nil
}

// getLookupCodes returns a list of LookupCodes for this element
func (WidgetCheckbox) getLookupCodes(element *Element, schemaElement schema.Element, lookupProvider LookupProvider) []LookupCode {

	lookupCodes := GetLookupCodes(element, schemaElement, lookupProvider)

	if len(lookupCodes) == 0 {
		lookupCodes = []LookupCode{
			{Value: "true", Label: element.Label},
		}
	}

	return lookupCodes
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (WidgetCheckbox) ShowLabels() bool {
	return false
}
