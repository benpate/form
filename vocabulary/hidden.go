package vocabulary

import (
	"github.com/benpate/convert"
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/schema"
)

// Hidden registers a text <input> widget into the library
func Hidden(library *form.Library) {

	library.Register("hidden", func(f form.Form, s *schema.Schema, v interface{}, b *html.Builder) error {

		// find the path and schema to use
		value, _, _ := s.Get(v, f.Path)
		valueString := convert.String(value)

		// Start building a new tag
		b.Input("hidden", f.Path).
			ID(f.ID).
			Value(valueString).
			Close()

		return nil
	})
}
