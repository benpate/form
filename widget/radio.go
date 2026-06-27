package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

// Radio is a widget that displays a list of radio buttons for selecting a single value from a list.
type Radio struct{}

// View is a part of the Widget interface.
// It builds the HTML for viewing this element.
func (widget Radio) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)
	valueString := e.GetString(value, &f.Schema)
	lookupCodes, _ := form.GetLookupCodes(e, schemaElement, provider)

	// Start building a new tag
	b.Div().Class("layout-value")
	for _, lookupCode := range lookupCodes {
		if lookupCode.Value == valueString {
			b.WriteString(lookupCode.Label)
			break
		}
	}
	b.Close()

	return nil
}

// Edit is a part of the Widget interface.
// It builds the HTML for editing this element.
func (widget Radio) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	// Calculate the element's ID
	id := e.ID

	if id == "" {
		id = e.Path + "." + e.Type
	}

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)
	valueString := e.GetString(value, &f.Schema)
	lookupCodes, _ := form.GetLookupCodes(e, schemaElement, provider)

	// Start building a new tag
	for _, lookupCode := range lookupCodes {
		radioID := id + "." + lookupCode.Value

		label := b.Label(radioID).
			ID(id + ".label")

		radio := b.Input("radio", e.Path).
			ID(radioID).
			Value(lookupCode.Value).
			Aria("label", lookupCode.Label).
			Aria("description", lookupCode.Description).
			TabIndex("0")

		if lookupCode.Value == valueString {
			radio.Attr("checked", "true")
		}

		radio.Close()

		b.Span().Aria("hidden", "true").InnerText(lookupCode.Label).Close()
		label.Close()
	}

	return nil
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For Radio widgets, labels are shown, so this always returns TRUE.
func (widget Radio) ShowLabels() bool {
	return true
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget Radio) ShowDescriptions() string {
	return "BOTTOM"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For Radio widgets, there is no special encoding,
// so this always returns an empty string.
func (widget Radio) Encoding(_ *form.Element) string {
	return ""
}
