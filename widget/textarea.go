package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

// TextArea renders a long text <textarea> widget
type TextArea struct{}

func (widget TextArea) View(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := element.GetString(value, s)

	// TODO: LOW: apply schema formats?
	b.Div().Class("layout-value").InnerText(valueString).Close()
	return nil
}

func (widget TextArea) Edit(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return TextArea{}.View(element, s, nil, value, b)
	}

	// find the path and schema to use
	schemaElement := element.GetSchema(s)
	valueString := element.GetString(value, s)

	if element.ID == "" {
		element.ID = element.Path + "." + element.Type
	}

	// Start building a new tag
	tag := b.Container("textarea").
		Name(element.Path).
		ID(element.ID).
		Attr("hint", element.Description).
		Attr("rows", element.Options.GetString("rows")).
		Aria("labelledby", element.ID+".label").
		Aria("describedby", element.ID+".description").
		TabIndex("0")

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
	tag.InnerText(valueString).Close()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget TextArea) ShowLabels() bool {
	return true
}

func (widget TextArea) Encoding(_ *form.Element) string {
	return ""
}
