package vocabulary

import (
	"github.com/benpate/convert"
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/schema"
)

// Checkbox registers a <input type="checkbox"> widget into the library
func Checkbox(library *form.Library) {

	library.Register("checkbox", func(f form.Form, s *schema.Schema, v interface{}, b *html.Builder) error {

		// find the path and schema to use
		_, value := locateSchema(f.Path, s, v)

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
