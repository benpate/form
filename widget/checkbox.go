package widget

import (
	"strings"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
)

// Checkbox is a widget that creates a set of checkbox input fields.
type Checkbox struct{}

// View is a part of the Widget interface.
// It builds the HTML for viewing this element.
func (widget Checkbox) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

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

// Edit is a part of the Widget interface.
// It builds the HTML for editing this element.
func (widget Checkbox) Edit(form *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := e.GetSchema(&form.Schema)
	valueSlice := e.GetSliceOfString(value, &form.Schema)
	lookupCodes := widget.getLookupCodes(e, schemaElement, provider)

	// Start building a new tag
	for _, lookupCode := range lookupCodes {
		id := "checkbox-" + strings.ReplaceAll(e.Path, ".", "-") + "-" + lookupCode.Value

		label := b.Label(id).ID("label-" + id)

		checkbox := b.Input("checkbox", e.Path).
			ID(id).
			Value(lookupCode.Value).
			Aria("label", lookupCode.Label).
			Aria("description", lookupCode.Description).
			TabIndex("0")

		if slice.Contains(valueSlice, lookupCode.Value) {
			checkbox.Attr("checked", "true")
		}
		checkbox.Close()

		b.Span().Aria("hidden", "true").InnerText(lookupCode.Label).Close()
		label.Close()
	}

	return nil
}

// getLookupCodes returns a list of LookupCodes for this element
func (widget Checkbox) getLookupCodes(element *form.Element, schemaElement schema.Element, provider form.LookupProvider) []form.LookupCode {

	lookupCodes, _ := form.GetLookupCodes(element, schemaElement, provider)

	if len(lookupCodes) == 0 {
		lookupCodes = []form.LookupCode{
			{Value: "true", Label: element.Label},
		}
	}

	return lookupCodes
}

/***********************************
 * Wiget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For Checkbox widgets, labels are not shown, so this always returns FALSE.
func (widget Checkbox) ShowLabels() bool {
	return false
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For Checkbox widgets, there is no special encoding,
// so this always returns an empty string.
func (widget Checkbox) Encoding(_ *form.Element) string {
	return ""
}
