package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("textarea", WidgetTextArea{})
}

// WidgetTextArea renders a long text <textarea> widget
type WidgetTextArea struct{}

func (WidgetTextArea) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := element.GetString(value, s)

	// TODO: LOW: apply schema formats?
	b.Div().Class("layout-value").InnerHTML(valueString).Close()
	return nil
}

func (WidgetTextArea) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetTextArea{}.View(element, s, lookupProvider, value, b)
	}

	// find the path and schema to use
	schemaElement := element.getElement(s)
	valueString := element.GetString(value, s)

	elementID := element.ID

	if elementID == "" {
		elementID = "textarea-" + element.Path
	}

	// Start building a new tag
	tag := b.Container("textarea").
		Name(element.Path).
		ID(elementID).
		Attr("hint", element.Description).
		Attr("rows", element.Options.GetString("rows"))

	if focus, ok := element.Options.GetBoolOK("focus"); ok && focus {
		tag.Attr("autofocus", "true")
	}

	// Add attributes that depend on what KIND of input we have.
	if schemaString, ok := schemaElement.(schema.String); ok {

		if schemaString.MinLength > 0 {
			tag.Attr("minlength", convert.String(schemaString.MinLength))
		}

		if schemaString.MaxLength > 0 {
			tag.Attr("maxlength", convert.String(schemaString.MaxLength))
		}

		if schemaString.Pattern != "" {
			tag.Attr("pattern", schemaString.Pattern)
		}

		if schemaString.Required {
			tag.Attr("required", "true")
		}
	}

	tag.TabIndex("0")
	tag.InnerHTML(valueString).Close()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (WidgetTextArea) ShowLabels() bool {
	return true
}
