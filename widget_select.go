package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("select", WidgetSelect{})
}

// WidgetSelect renders a select box widget
type WidgetSelect struct{}

func (WidgetSelect) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := element.getElement(s)
	valueString := element.GetString(value, s)
	lookupCodes := GetLookupCodes(element, schemaElement, lookupProvider)

	// Start building a new tag
	b.Div().Class("layout-value").EndBracket()
	for _, lookupCode := range lookupCodes {
		if lookupCode.Value == valueString {
			b.WriteString(lookupCode.Label)
			break
		}
	}

	// TODO: HIGH: Add Support for WritableLookupProvider

	b.Close()

	return nil
}

func (WidgetSelect) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetSelect{}.View(element, s, lookupProvider, value, b)
	}

	_, isWritable := lookupProvider.(WritableLookupGroup)

	// find the path and schema to use
	schemaElement := element.getElement(s)
	valueString := element.GetString(value, s)

	if element, ok := schemaElement.(schema.Array); ok {
		schemaElement = element.Items
	}

	// Get all lookupCodes for this element...
	lookupCodes := GetLookupCodes(element, schemaElement, lookupProvider)

	elementID := element.ID

	if elementID == "" {
		elementID = "select-" + element.Path
	}

	selectBox := b.Container("select").
		ID(elementID).
		Name(element.Path).
		TabIndex("0")

	if element.Options.GetBool("focus") {
		selectBox.Attr("autofocus", "true")
	}

	// Calculate scripts
	if isWritable {
		selectBox.Script(`on change if my value is '::NEWVALUE::' then 
			set newName to window.prompt("Add New Value")
			if newName is not null then 
				make an Option from newName, ("::NEWVALUE::" + newName), true, true called newOption 
				call me.add(newOption)
			end`)
	}

	if (schemaElement != nil) && (!schemaElement.IsRequired()) {
		b.Container("option").Value("").InnerHTML("").Close()
	}

	for _, lookupCode := range lookupCodes {
		opt := b.Container("option").Value(lookupCode.Value)
		if lookupCode.Value == valueString {
			opt.Attr("selected", "true")
		}
		opt.InnerHTML(lookupCode.Label).Close()
	}

	if isWritable {
		b.Container("option").
			Class("add-new").
			Script("on change log me").
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

func (WidgetSelect) ShowLabels() bool {
	return true
}
