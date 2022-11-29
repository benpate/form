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
	b.Close()

	return nil
}

func (WidgetSelect) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetSelect{}.View(element, s, lookupProvider, value, b)
	}

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

	if optionURL := element.Options.GetString("optionUrl"); optionURL != "" {
		selectBox.Script("on load fetch " + optionURL + " as json then set options to it.map(\\x -> new Option(x.Label, x.Value)) the set my.options to options")
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

	b.CloseAll()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (WidgetSelect) ShowLabels() bool {
	return true
}
