package vocabulary

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

// Option registers a text <input> widget into the library
func Option(library *form.Library) {

	library.Register("option", func(f form.Form, s *schema.Schema, v interface{}, b *html.Builder) error {

		// find the path and schema to use
		value, schemaElement, _ := s.Get(v, f.Path)
		valueString := convert.String(value)
		format := convert.String(f.Options["format"])

		if format == "" {

			switch schemaElement.(type) {
			case schema.Array:
				format = "checkbox"
			default:
				format = "radio"
			}
		}

		// Start building a new tag

		b.Label(f.ID)

		b.Input(format, f.Path).
			ID(f.ID).
			Value(valueString).
			Class(f.CSSClass).
			InnerHTML(f.Label)

		b.CloseAll()

		return nil
	})
}
