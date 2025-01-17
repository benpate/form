package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

type Password struct{}

func (widget Password) View(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {
	b.Div().Class("layout-value", element.Options.GetString("class")).InnerText("********").Close()
	return nil
}

func (widget Password) Edit(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return Password{}.View(element, s, nil, value, b)
	}

	// find the path and schema to use
	schemaElement := element.GetSchema(s)

	if element.ID == "" {
		element.ID = element.Path + "." + element.Type
	}

	// Start building a new tag
	tag := b.Input("password", element.Path).
		ID(element.ID).
		Aria("label", element.Label).
		Aria("description", element.Description).
		TabIndex("0")

	if focus, ok := element.Options.GetBoolOK("focus"); ok && focus {
		tag.Attr("autofocus", "true")
	}

	if placeholder := element.Options.GetString("placeholder"); placeholder != "" {
		tag.Attr("placeholder", placeholder)
	}

	// Custom CSS style
	if style := element.Options.GetString("style"); style != "" {
		tag.Attr("style", style)
	}

	// Add attributes that depend on what KIND of input we have.
	switch s := schemaElement.(type) {

	case schema.String:

		if s.MinLength > 0 {
			tag.Attr("minlength", convert.String(s.MinLength))
		}

		if s.MaxLength > 0 {
			tag.Attr("maxlength", convert.String(s.MaxLength))
		}

		if s.Pattern != "" {
			tag.Attr("pattern", s.Pattern)
		}

		if s.Required {
			tag.Attr("required", "true")
		}
	}

	if autocomplete := element.Options.GetString("autocomplete"); autocomplete != "" {
		tag.Attr("autocomplete", autocomplete)

		if autocomplete == "off" {
			tag.Attr("data-1p-ignore", "true")
		}
	}

	b.CloseAll()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget Password) ShowLabels() bool {
	return true
}

func (widget Password) Encoding(_ *form.Element) string {
	return ""
}

/*


   <input type="button">
   <input type="checkbox">
   <input type="color">
   <input type="date">
   <input type="datetime-local">
   <input type="email">
   <input type="file">
   <input type="hidden">
   <input type="image">
   <input type="month">
   <input type="number">
   <input type="password">
   <input type="radio">
   <input type="range">
   <input type="reset">
   <input type="search">
   <input type="submit">
   <input type="tel">
   <input type="text">
   <input type="time">
   <input type="url">
   <input type="week">
*/
