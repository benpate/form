package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("toggle", HTMLToggle)
}

// HTMLToggle registers a custom toggle widget into the library
func HTMLToggle(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString, _ := element.GetString(value, s)

	// Start building a new tag
	tag := b.Span().Script("install toggle").Name(element.Path)

	if convert.Bool(valueString) {
		tag.Value("true")
	}

	tag.Attr("true-text", element.Options.GetString("true-text"))
	tag.Attr("false-text", element.Options.GetString("false-text"))

	b.CloseAll()
	return nil
}
