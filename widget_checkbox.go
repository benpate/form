package form

import (
	"strings"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
)

func init() {
	Register("checkbox", HTMLCheckbox)
}

func HTMLCheckbox(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueSlice, schemaElement := element.GetSliceOfString(value, schema)

	lookupCodes := GetLookupCodes(element, schemaElement, lookupProvider)

	if len(lookupCodes) == 0 {
		lookupCodes = []LookupCode{
			{Value: "true", Label: element.Label},
		}
	}

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
