package vocabulary

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/compare"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

// Multiselect registers a custom multi-select widget into the library
func Multiselect(library *form.Library) {

	library.Register("multiselect", func(f form.Form, s *schema.Schema, v interface{}, b *html.Builder) error {

		// find the path and schema to use
		value, schemaElement, _ := s.Get(v, f.Path)

		if element, ok := schemaElement.(schema.Array); ok {
			schemaElement = element.Items
		}

		sortable := convert.Bool(f.Options["sort"])

		maxHeight := "300"

		if value, ok := f.Options["maxHeight"].(string); ok {
			maxHeight = value
		}

		// Get all options for this element...
		options := library.Options(f, schemaElement)
		valueSlice := convert.SliceOfString(value)

		b.Div().Class("multiselect").Script("install multiselect(sort:" + convert.String(sortable) + ")")
		b.Div().Class("options").Style("maxHeight:" + maxHeight + "px")

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

			b.Close() // input

			b.Div()
			b.Div().InnerHTML(option.Label).Close()

			if option.Description != "" {
				b.Div().Class("text-sm", "gray50").InnerHTML(option.Description).Close()
			}
			b.Close() // div
			b.Close() // label
		}

		b.Close() // .options

		// Buttons
		if sortable {
			b.Div().Class("buttons").EndBracket()
			b.Button().Type("button").Data("sort", "up").InnerHTML("△").Close()
			b.Button().Type("button").Data("sort", "down").InnerHTML("▽").Close()
			b.Close()
		}

		b.CloseAll()
		return nil
	})
}
