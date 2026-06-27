package widget

import (
	"strings"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
)

// RadioButtonGroup renders a group of fancy radio buttons that look like buttons
type RadioButtonGroup struct{}

// View generates the read-only HTML for this radio-button group, showing the selected label.
func (widget RadioButtonGroup) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the schema and value to use
	schemaElement := e.GetSchema(&f.Schema)
	valueSlice := e.GetSliceOfString(value, &f.Schema)
	lookupCodes := widget.getLookupCodes(e, schemaElement, provider)

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

// Edit generates the editable HTML for this radio-button group.
func (widget RadioButtonGroup) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)
	valueSlice := e.GetSliceOfString(value, &f.Schema)
	lookupCodes := widget.getLookupCodes(e, schemaElement, provider)

	// Start building a new tag
	for _, lookupCode := range lookupCodes {
		id := "radiobutton-" + strings.ReplaceAll(e.Path, ".", "-") + "-" + lookupCode.Value

		b.Label(id).ID("label-" + id).Class("radiobutton")

		if lookupCode.Icon != "" {
			b.I().Class("margin-horizontal", "bi", "bi-"+lookupCode.Icon).Style("font-size:32px;").Close()
		}

		b.Div().Class("flex-column")
		b.Div().Class("bold").InnerText(lookupCode.Label).Close()
		b.Div().Class("text-sm", "text-gray").InnerText(lookupCode.Description).Close()

		toggleButton := b.Input("radio", e.Path)

		toggleButton.
			ID(id).
			Class(e.Options.GetString("class")).
			Value(lookupCode.Value).
			Aria("label", lookupCode.Label).
			Aria("description", lookupCode.Description).
			Script(e.Options.GetString("script")).
			TabIndex("0")

		if slice.Contains(valueSlice, lookupCode.Value) {
			toggleButton.Attr("checked", "true")
		}

		b.Close()
		b.Close()
		b.Close()
	}

	return nil
}

// getLookupCodes returns a list of LookupCodes for this element
func (widget RadioButtonGroup) getLookupCodes(e *form.Element, schemaElement schema.Element, provider form.LookupProvider) []form.LookupCode {

	lookupCodes, _ := form.GetLookupCodes(e, schemaElement, provider)

	if len(lookupCodes) == 0 {
		lookupCodes = []form.LookupCode{
			{Value: "true", Label: e.Label},
		}
	}

	return lookupCodes
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For RadioButtonGroup widgets, labels are shown, so this always returns TRUE.
func (widget RadioButtonGroup) ShowLabels() bool {
	return true
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget RadioButtonGroup) ShowDescriptions() string {
	return "TOP"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For RadioButtonGroup widgets, there is no special encoding,
// so this always returns an empty string.
func (widget RadioButtonGroup) Encoding(_ *form.Element) string {
	return ""
}
