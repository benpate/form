package vocabulary

import (
	"github.com/benpate/compare"
	"github.com/benpate/convert"
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/schema"
)

// Multiselect registers a custom multi-select widget into the library
func Multiselect(library *form.Library) {

	library.Register("multiselect", func(f form.Form, s *schema.Schema, v interface{}, b *html.Builder) error {

		// find the path and schema to use
		schemaElement, value := locateSchema(f.Path, s, v)

		if element, ok := schemaElement.(schema.Array); ok {
			schemaElement = element.Items
		}

		maxHeight := "200"

		if value, ok := f.Options["maxHeight"].(string); ok {
			maxHeight = value
		}

		// Get all options for this element...
		options := library.Options(f, schemaElement)

		valueSlice := convert.SliceOfString(value)

		b.Div().Class("multiselect").Style("maxHeight:" + maxHeight + "px").Script("install multiselect")

		for _, option := range options {

			var id string

			if f.ID != "" {
				id = f.ID + "_" + option.Value
			} else {
				id = f.Path + "_" + option.Value
			}

			b.Label(id)

			input := b.Input("checkbox", f.Path).ID(id).Value(option.Value)

			if compare.Contains(valueSlice, option.Value) {
				input.Attr("checked", "true")
			}

			b.Close()

			b.Div()
			b.Div().InnerHTML(option.Label).Close()

			if option.Description != "" {
				b.Div().Class("text-sm", "gray50").InnerHTML(option.Description).Close()
			}

			b.Close()
			b.Close()
		}

		b.CloseAll()
		return nil
	})
}
