package form

import (
	"github.com/benpate/html"
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

	// TODO: apply schema formats?
	b.Div().Class("layout-value").InnerHTML(valueString).Close()
	return nil
}

func (WidgetTextArea) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := element.getElement(s)
	valueString := element.GetString(value, s)

	// Start building a new tag
	tag := b.Container("textarea").
		Name(element.Path).
		ID(element.ID).
		Attr("hint", element.Description).
		Attr("rows", element.Options.GetString("rows"))

	// Add attributes that depend on what KIND of input we have.
	if schemaString, ok := schemaElement.(schema.String); ok {

		if schemaString.MinLength.IsPresent() {
			tag.Attr("minlength", schemaString.MinLength.String())
		}

		if schemaString.MaxLength.IsPresent() {
			tag.Attr("maxlength", schemaString.MaxLength.String())
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
