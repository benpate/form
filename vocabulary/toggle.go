package vocabulary

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

// Toggle registers a custom toggle widget into the library
func Toggle(library *form.Library) {

	library.Register("toggle", func(f form.Form, s *schema.Schema, v interface{}, b *html.Builder) error {

		// find the path and schema to use
		value, _, _ := s.Get(v, f.Path)

		// Start building a new tag
		tag := b.Span().Script("install toggle").Name(f.Path)

		if convert.Bool(value) {
			tag.Value("true")
		}

		tag.Attr("true-text", convert.String(f.Options["true-text"]))
		tag.Attr("false-text", convert.String(f.Options["false-text"]))

		b.CloseAll()
		return nil
	})
}
