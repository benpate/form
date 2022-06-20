package vocabulary

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/compare"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

// Select registers a text <input> widget into the library
func Select(library *form.Library) {

	library.Register("select", func(f form.Form, s *schema.Schema, v interface{}, b *html.Builder) error {

		var selectMany bool

		// find the path and schema to use
		value, schemaElement, _ := s.Get(v, f.Path)

		if element, ok := schemaElement.(schema.Array); ok {
			schemaElement = element.Items
			selectMany = true
		}

		// Get all options for this element...
		options := library.Options(f, schemaElement)

		// SelectMany
		if selectMany {

			valueSlice := convert.SliceOfString(value)

			for _, option := range options {
				label := b.Label(f.ID)

				input := b.Input("checkbox", f.Path).ID(f.ID).Value(option.Value)

				if compare.Contains(valueSlice, option.Value) {
					input.Attr("checked", "true")
				}

				input.TabIndex("0")
				input.Close()
				label.InnerHTML(option.Label)
				label.Close()
			}

			b.CloseAll()
			return nil
		}

		// SelectOne
		valueString := convert.String(value)

		if f.Options["format"] == "radio" {

			for _, option := range options {
				label := b.Label(f.ID)

				input := b.Input("radio", f.Path).
					ID(f.ID).
					Value(option.Value)

				if valueString == option.Value {
					input.Attr("checked", "true")
				}

				input.TabIndex("0")
				input.Close()
				label.InnerHTML(option.Label)
				label.Close()
			}

		} else {

			// Fall through to select box

			dropdown := b.Container("select").ID(f.ID).Name(f.Path).Class(f.CSSClass).TabIndex("0")

			if !schemaElement.IsRequired() {
				b.Container("option").Value("").InnerHTML("").Close()
			}

			for _, option := range options {
				opt := b.Container("option").Value(option.Value)
				if option.Value == valueString {
					opt.Attr("selected", "true")
				}
				opt.InnerHTML(option.Label).Close()
			}
			dropdown.Close()
		}

		b.CloseAll()
		return nil
	})
}
