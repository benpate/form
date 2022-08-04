package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/compare"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/first"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("multiselect", HTMLMultiselect)
}

// Multiselect registers a custom multi-select widget into the library
func HTMLMultiselect(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueSlice, schemaElement := element.GetSliceOfString(value, s)

	sortable := element.Options.GetBool("sort")
	maxHeight := first.String(element.Options.GetString("maxHeight"), "300")

	// Get all options for this element...
	options := GetLookupCodes(element, schemaElement, lookupProvider)

	b.Div().Class("multiselect").Script("install multiselect(sort:" + convert.String(sortable) + ")")
	b.Div().Class("options").Style("maxHeight:" + maxHeight + "px")

	for _, option := range options {
		id := element.Path + "." + option.Value

		b.Label(id)

		input := b.Input("checkbox", element.Path).ID(id).Value(option.Value)

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
}
