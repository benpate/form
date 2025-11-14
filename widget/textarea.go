package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

// TextArea renders a long text <textarea> widget
type TextArea struct{}

func (widget TextArea) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)

	// TODO: LOW: apply schema formats?
	b.Div().Class("layout-value").InnerText(valueString).Close()
	return nil
}

func (widget TextArea) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)
	valueString := e.GetString(value, &f.Schema)

	if e.ID == "" {
		e.ID = e.Path + "." + e.Type
	}

	// Start building a new tag
	tag := b.Container("textarea").
		Name(e.Path).
		ID(e.ID).
		Attr("hint", e.Description).
		Attr("rows", e.Options.GetString("rows")).
		Aria("labelledby", e.ID+".label").
		Aria("describedby", e.ID+".description").
		TabIndex("0")

	// Autofocus
	if focus, ok := e.Options.GetBoolOK("focus"); ok && focus {
		tag.Attr("autofocus", "true")
	}

	// Custom CSS style
	if style := e.Options.GetString("style"); style != "" {
		tag.Attr("style", style)
	}

	// Add placeholder
	if placeholder := e.Options.GetString("placeholder"); placeholder != "" {
		tag.Attr("placeholder", placeholder)
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

		if schemaString.MaxLength > 0 {
			if e.Options.GetBool("showLimit") {
				tag.Attr("script", "install showLimit")
			}
		}
	}

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
