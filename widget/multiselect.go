package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/compare"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/first"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
)

type Multiselect struct{}

func (widget Multiselect) View(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	schemaElement := element.GetSchema(s)
	valueSlice := element.GetSliceOfString(value, s)
	lookupCodes, _ := form.GetLookupCodes(element, schemaElement, provider)
	first := true

	b.Div().Class("layout-value")
	for _, lookupCode := range lookupCodes {

		if slice.Contains(valueSlice, lookupCode.Value) {

			if first {
				first = false
			} else {
				b.WriteString(", ")
			}

			b.WriteString(lookupCode.Label)
		}
	}
	b.Close()

	return nil
}

// Multiselect registers a custom multi-select widget into the library
func (widget Multiselect) Edit(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return Multiselect{}.View(element, s, provider, value, b)
	}

	// find the path and schema to use
	schemaElement := element.GetSchema(s)
	valueSlice := element.GetSliceOfString(value, s)

	sortable, _ := element.Options.GetBoolOK("sort")
	maxHeight := first.String(element.Options.GetString("maxHeight"), "300")

	// Get all options for this element...
	options, _ := form.GetLookupCodes(element, schemaElement, provider)

	b.Div().Class("multiselect").Script("install multiselect(sort:" + convert.String(sortable) + ")")
	b.Div().Class("options").Style("max-height:" + maxHeight + "px")

	elementID := element.ID

	if elementID == "" {
		elementID = "multiselect-" + element.Path
	}

	for _, option := range options {

		optionID := elementID + "-" + option.Value

		b.Label(optionID)

		input := b.Input("checkbox", element.Path).ID(optionID).Value(option.Value)

		if compare.Contains(valueSlice, option.Value) {
			input.Attr("checked", "true")
		}

		b.Close() // input

		b.Div()
		b.Div().InnerText(option.Label).Close()

		if option.Description != "" {
			b.Div().Class("text-sm", "gray50").InnerText(option.Description).Close()
		}
		b.Close() // div
		b.Close() // label
	}

	// TODO: LOW: Add support for WritableLookupProvider

	b.Close() // .options

	// Buttons
	if sortable {
		b.Div().Class("buttons").EndBracket()
		b.Button().Type("button").Data("sort", "up").InnerText("△").Close()
		b.Button().Type("button").Data("sort", "down").InnerText("▽").Close()
		b.Close()
	}

	b.CloseAll()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget Multiselect) ShowLabels() bool {
	return true
}

func (widget Multiselect) Encoding(_ *form.Element) string {
	return ""
}
