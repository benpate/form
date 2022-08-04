package form

import (
	"strings"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("textarea", HTMLTextArea)
}

// HTMLTextarea registers a <textarea> input widget into the library
func HTMLTextArea(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString, schemaElement := element.GetString(value, s)
	id := "textarea-" + strings.ReplaceAll(element.Path, ".", "-")

	// Start building a new tag
	tag := b.Container("textarea").
		Name(element.Path).
		ID(id).
		Attr("hint", element.Description).
		Attr("rows", element.Options.GetString("rows"))

	// Add attributes that depend on what KIND of input we have.
	if schemaString, ok := schemaElement.(schema.String); ok {

		if schemaString.MinLength.IsPresent() {
			tag.Attr("minlength", schemaString.MinLength.String())
		}

		if schemaString.MaxLength.IsPresent() {
			tag.Attr("maxlength", schemaString.MaxLength.String())
		}

		if schemaString.Pattern != "" {
			tag.Attr("pattern", schemaString.Pattern)
		}

		if schemaString.Required {
			tag.Attr("required", "true")
		}
	}

	tag.TabIndex("0")
	tag.InnerHTML(valueString).Close()
	return nil
}
