package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/form/groupie"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/compare"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/first"
	"github.com/benpate/rosetta/slice"
)

type Multiselect struct{}

func (widget Multiselect) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	schemaElement := e.GetSchema(&f.Schema)
	valueSlice := e.GetSliceOfString(value, &f.Schema)
	lookupCodes, _ := form.GetLookupCodes(e, schemaElement, provider)
	first := true

	group := groupie.New()

	b.Div().Class("layout-value")
	for _, lookupCode := range lookupCodes {

		if group.Header(lookupCode.Group) {
			b.WriteString(lookupCode.Group + ": ")
		}

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
func (widget Multiselect) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	if e.ReadOnly {
		return Multiselect{}.View(f, e, provider, value, b)
	}

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)

	sortable, _ := e.Options.GetBoolOK("sort")
	maxHeight := first.String(e.Options.GetString("maxHeight"), "300")

	// Get all options for this element...
	options, _ := form.GetLookupCodes(e, schemaElement, provider)

	b.Div().Class("multiselect").Script("install multiselect(sort:" + convert.String(sortable) + ")")
	b.Div().Class("options").Style("max-height:" + maxHeight + "px")

	elementID := e.ID

	if elementID == "" {
		elementID = "multiselect-" + e.Path
	}

	group := groupie.New()

	for _, option := range options {

		if group.Header(option.Group) {
			if option.Group != "" {
				b.Div().Class("multiselect-header").InnerText(option.Group).Close()
			}
		}

		optionID := elementID + "-" + option.Value

		b.Label(optionID)

		input := b.Input("checkbox", e.Path).ID(optionID).Value(option.Value) // nolint:scopeguard b.Input has a side effect

		if valueSlice := e.GetSliceOfString(value, &f.Schema); compare.Contains(valueSlice, option.Value) {
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
