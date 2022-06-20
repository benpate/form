package vocabulary

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

// Checkbox registers a <input type="checkbox"> widget into the library
func Checkbox(library *form.Library) {

	library.Register("checkbox", func(f form.Form, s *schema.Schema, v interface{}, b *html.Builder) error {

		// find the path and schema to use
		value, _, _ := s.Get(v, f.Path)

		// Start building a new tag
		tag := b.Input("checkbox", f.Path).
			ID(f.ID).
			Style("vertical-align:text-bottom").
			Value("true")

		if convert.Bool(value) {
			tag.Attr("checked", "true")
		}

		if f.CSSClass != "" {
			tag.Attr("class", f.CSSClass)
		}

		if f.Description != "" {
			tag.Attr("hint", f.Description)
		}

		tag.TabIndex("0")
		b.CloseAll()
		return nil
	})
}
