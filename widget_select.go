package form

import (
	"strings"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("select", HTMLSelect)
}

// HTMLSelect registers a select box widget into the library
func HTMLSelect(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString, schemaElement := element.GetString(value, s)
	id := "select-" + strings.ReplaceAll(element.Path, ".", "-")

	if element, ok := schemaElement.(schema.Array); ok {
		schemaElement = element.Items
	}

	// Get all lookupCodes for this element...
	lookupCodes := GetLookupCodes(element, schemaElement, lookupProvider)

	b.Container("select").
		ID(id).
		Name(element.Path).
		TabIndex("0")

	if !schemaElement.IsRequired() {
		b.Container("option").Value("").InnerHTML("").Close()
	}

	for _, lookupCode := range lookupCodes {
		opt := b.Container("option").Value(lookupCode.Value)
		if lookupCode.Value == valueString {
			opt.Attr("selected", "true")
		}
		opt.InnerHTML(lookupCode.Label).Close()
	}

	b.CloseAll()
	return nil
}
