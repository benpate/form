package vocabulary

import (
	"github.com/benpate/convert"
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/schema"
)

// Textarea registers a <textarea> input widget into the library
func Textarea(library *form.Library) {

	library.Register("textarea", func(f form.Form, s *schema.Schema, v interface{}, b *html.Builder) error {

		// find the path and schema to use
		value, schemaElement, _ := s.Get(v, f.Path)

		valueString := convert.String(value)

		// Start building a new tag
		tag := b.Container("textarea").
			ID(f.ID).
			Name(f.Path).
			Class(f.CSSClass).
			Attr("hint", f.Description)

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
	})
}
