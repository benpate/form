package widget

import (
	"strings"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

type Radio struct{}

func (widget Radio) View(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := element.GetSchema(s)
	valueString := element.GetString(value, s)
	lookupCodes, _ := form.GetLookupCodes(element, schemaElement, provider)

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

func (widget Radio) Edit(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return Radio{}.View(element, s, provider, value, b)
	}

	// Calculate the element's ID
	id := element.ID

	if id == "" {
		id = "radio-" + strings.ReplaceAll(element.Path, ".", "-")
	}

	// find the path and schema to use
	schemaElement := element.GetSchema(s)
	valueString := element.GetString(value, s)
	lookupCodes, _ := form.GetLookupCodes(element, schemaElement, provider)

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

func (widget Radio) ShowLabels() bool {
	return true
}

func (widget Radio) Encoding(_ *form.Element) string {
	return ""
}
