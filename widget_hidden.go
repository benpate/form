package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("hidden", HTMLHidden)
}

// Hidden registers a text <input> widget into the library
func HTMLHidden(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	elementValue, _ := element.GetString(value, schema)

	// Start building a new tag
	b.Input("hidden", element.Path).
		ID(element.Path).
		Value(elementValue).
		Close()

	return nil
}
