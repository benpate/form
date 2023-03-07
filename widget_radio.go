package form

import (
	"strings"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("radio", WidgetRadio{})
}

type WidgetRadio struct{}

func (WidgetRadio) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := element.getElement(s)
	valueString := element.GetString(value, s)
	lookupCodes, _ := GetLookupCodes(element, schemaElement, lookupProvider)

	// Start building a new tag
	b.Div().Class("layout-value")
	for _, lookupCode := range lookupCodes {
		if lookupCode.Value == valueString {
			b.WriteString(lookupCode.Label)
			break
		}
	}
	b.Close()

	return nil
}

func (WidgetRadio) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetRadio{}.View(element, s, lookupProvider, value, b)
	}

	// Calculate the element's ID
	id := element.ID

	if id == "" {
		id = "radio-" + strings.ReplaceAll(element.Path, ".", "-")
	}

	// find the path and schema to use
	schemaElement := element.getElement(s)
	valueString := element.GetString(value, s)
	lookupCodes, _ := GetLookupCodes(element, schemaElement, lookupProvider)

	// Start building a new tag
	for _, lookupCode := range lookupCodes {
		radioID := id + "-" + lookupCode.Value
		b.Label(radioID)

		radio := b.Input("radio", element.Path).
			ID(radioID).
			Value(lookupCode.Value)

		if lookupCode.Value == valueString {
			radio.Attr("checked", "true")
		}

		radio.InnerText(lookupCode.Label).Close()
		b.CloseAll()
	}

	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (WidgetRadio) ShowLabels() bool {
	return true
}
