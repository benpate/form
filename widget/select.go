package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
)

// Select renders a select box widget
type Select struct{}

func (widget Select) View(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := element.GetSchema(s)
	valueString := element.GetString(value, s)
	lookupCodes, _ := form.GetLookupCodes(element, schemaElement, provider)

	// Start building a new tag
	b.Div().Class("layout-value").EndBracket()
	for _, lookupCode := range lookupCodes {
		if lookupCode.Value == valueString {
			b.WriteString(lookupCode.Label)
			break
		}
	}

	b.Close()

	return nil
}

func (widget Select) Edit(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return Select{}.View(element, s, provider, value, b)
	}

	// find the path and schema to use
	schemaElement := element.GetSchema(s)
	valueSlice := element.GetSliceOfString(value, s)

	if element, ok := schemaElement.(schema.Array); ok {
		schemaElement = element.Items
	}

	// Get all lookupCodes for this element...
	lookupCodes, isWritable := form.GetLookupCodes(element, schemaElement, provider)

	elementID := element.ID

	if elementID == "" {
		elementID = "select-" + element.Path
	}

	selectBox := b.Container("select").
		ID(elementID).
		Name(element.Path).
		TabIndex("0")

	if focus, ok := element.Options.GetBoolOK("focus"); ok && focus {
		selectBox.Attr("autofocus", "true")
	}

	// Support for writable lookup providers
	if isWritable {
		selectBox.Script(`on change if my value is '::NEWVALUE::' then 
			set newName to window.prompt("Add New Value")
			if newName is not null then 
				make an Option from newName, ("::NEWVALUE::" + newName), true, true called newOption 
				call me.add(newOption)
			end`)
	}

	// Allow null options if not required
	if (schemaElement != nil) && (!schemaElement.IsRequired()) {
		b.Container("option").Value("").InnerText("").Close()
	}

	// Display all lookup codes
	for _, lookupCode := range lookupCodes {
		opt := b.Container("option").Value(lookupCode.Value)
		if slice.Contains(valueSlice, lookupCode.Value) {
			opt.Attr("selected", "true")
		}
		opt.InnerText(lookupCode.Label).Close()
	}

	// Support for writable lookup providers
	if isWritable {
		b.Container("option").
			Class("add-new").
			Value("::NEWVALUE::").
			InnerHTML("+ Add Another...").
			Close()
	}

	b.CloseAll()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget Select) ShowLabels() bool {
	return true
}

func (widget Select) Encoding(_ *form.Element) string {
	return ""
}
