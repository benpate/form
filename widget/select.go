package widget

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
)

func init() {
	Register("select", Select{})
}

// Select renders a select box widget
type Select struct{}

func (widget Select) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := element.getElement(s)
	valueString := element.GetString(value, s)
	lookupCodes, _ := GetLookupCodes(element, schemaElement, lookupProvider)

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

func (widget Select) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return Select{}.View(element, s, lookupProvider, value, b)
	}

	// find the path and schema to use
	schemaElement := element.getElement(s)
	valueSlice := element.GetSliceOfString(value, s)

	if element, ok := schemaElement.(schema.Array); ok {
		schemaElement = element.Items
	}

	// Get all lookupCodes for this element...
	lookupCodes, isWritable := GetLookupCodes(element, schemaElement, lookupProvider)

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

func (widget Select) Encoding(_ *Element) string {
	return ""
}
