package vocabulary

import (
	"github.com/benpate/builder"
	"github.com/benpate/form"
	"github.com/benpate/schema"
)

// Textarea registers a <textarea> input widget into the library
func Textarea(library form.Library) {

	library.Register("textarea", func(f form.Form, s *schema.Schema, v interface{}, b *builder.Builder) error {

		// find the path and schema to use
		schemaObject, valueString := locateSchema(f.Path, s, v)

		// Start building a new tag
		tag := b.Container("textarea").
			ID(f.ID).
			Name(f.Path).
			Class(f.CSSClass).
			Attr("hint", f.Description)

		// Add attributes that depend on what KIND of input we have.
		if schemaString, ok := schemaObject.(schema.String); ok {

			if schemaString.MinLength.IsPresent() {
				tag.Attr("minlength", schemaString.MinLength.Int())
			}

			if schemaString.MaxLength.IsPresent() {
				tag.Attr("maxlength", schemaString.MaxLength.Int())
			}

			if schemaString.Pattern != "" {
				tag.Attr("pattern", schemaString.Pattern)
			}

			if schemaString.Required {
				tag.Attr("required", true)
			}
		}

		tag.InnerHTML(valueString).Close()
		return nil
	})
}
