package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

type Radio struct{}

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
 * Wiget Metadata
 ***********************************/

func (widget Radio) ShowLabels() bool {
	return true
}

func (widget Radio) Encoding(_ *form.Element) string {
	return ""
}
