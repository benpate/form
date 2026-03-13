package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

// RadioColors is a widget that displays a list of radio buttons for selecting a single value from a list.
type RadioColors struct{}

// View is a part of the Widget interface.
// It builds the HTML for viewing this element.
func (widget RadioColors) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

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
func (widget RadioColors) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	// Calculate the element's ID
	id := e.ID

	if id == "" {
		id = e.Path + "." + e.Type
	}

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)
	valueString := e.GetString(value, &f.Schema)
	lookupCodes, _ := form.GetLookupCodes(e, schemaElement, provider)

	// Wrapper to display radio buttons in a row
	b.Span().Class("flex-row")

	// Start building a new tag
	for _, lookupCode := range lookupCodes {
		radioID := id + "." + lookupCode.Value

		label := b.Label(radioID).
			ID(id+".label").
			Class("flex-column", "flex-align-center")

		b.Div().Style("height: 36px", "width:32px", "border-radius:12px", "margin-bottom:8px", "background-color:"+lookupCode.Value).Close()

		radio := b.Input("radio", e.Path).
			ID(radioID).
			Value(lookupCode.Value).
			Aria("label", lookupCode.Label).
			Aria("description", lookupCode.Description).
			TabIndex("0").
			Style("margin:0", "padding:0")

		if lookupCode.Value == valueString {
			radio.Attr("checked", "true")
		}

		radio.Close()
		label.Close()
	}

	b.CloseAll()

	return nil
}

/***********************************
 * Widget Metadata
 ***********************************/

func (widget RadioColors) ShowLabels() bool {
	return true
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget RadioColors) ShowDescriptions() string {
	return "TOP"
}

func (widget RadioColors) Encoding(_ *form.Element) string {
	return ""
}
