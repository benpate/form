package form

import (
	"strings"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("radio", HTMLRadio)
}

func HTMLRadio(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString, schemaElement := element.GetString(value, schema)
	lookupCodes := GetLookupCodes(element, schemaElement, lookupProvider)

	// Start building a new tag
	for _, lookupCode := range lookupCodes {
		id := "radio-" + strings.ReplaceAll(element.Path, ".", "-") + "-" + lookupCode.Value
		b.Label(id)

		checkbox := b.Input("radio", element.Path).
			ID(id).
			Value(lookupCode.Value)

		if lookupCode.Value == valueString {
			checkbox.Attr("checked", "true")
		}

		checkbox.InnerHTML(lookupCode.Label).Close()
		b.CloseAll()
	}

	return nil
}
