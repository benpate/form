package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

type Text struct{}

func (widget Text) View(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := element.GetString(value, s)

	// TODO: LOW: Apply formatting options?
	b.Div().Class("layout-value", element.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

func (widget Text) Edit(element *form.Element, s *schema.Schema, provider form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return Text{}.View(element, s, nil, value, b)
	}

	// find the path and schema to use
	schemaElement := element.GetSchema(s)
	valueString := element.GetString(value, s)

	if element.ID == "" {
		element.ID = element.Path + "." + element.Type
	}

	// Start building a new tag
	tag := b.Input("", element.Path).
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

	// Enumeration Options
	lookupCodes, _ := form.GetLookupCodes(element, schemaElement, provider)
	if len(lookupCodes) > 0 {
		tag.Attr("list", "datalist-"+element.Path)
	}

	// Add attributes that depend on what KIND of input we have.
	switch s := schemaElement.(type) {

	case schema.Integer:

		tag.Type("number")

		tag.Attr("step", convert.String(convert.IntDefault(element.Options["step"], 1)))

		if s.Minimum.IsPresent() {
			tag.Attr("min", s.Minimum.String())
		}

		if s.Maximum.IsPresent() {
			tag.Attr("max", s.Maximum.String())
		}

		if s.Required {
			tag.Attr("required", "true")
		}

		if s.RequiredIf != "" {
			tag.Script("install requiredIf(condition:'" + s.RequiredIf + "')")
		}

	case schema.Number:

		tag.Type("number")

		tag.Attr("step", convert.String(convert.FloatDefault(element.Options["step"], 0.01)))

		if s.Minimum.IsPresent() {
			tag.Attr("min", s.Minimum.String())
		}

		if s.Maximum.IsPresent() {
			tag.Attr("max", s.Maximum.String())
		}

		if s.Required {
			tag.Attr("required", "true")
		}

	case schema.String:

		switch s.Format {

		case "date":
			tag.Type("date")

		case "datetime":
			tag.Type("datetime-local")

		case "email":
			tag.Type("email")

		case "time":
			tag.Type("time")

		case "tel":
			tag.Type("tel")

		case "url":
			tag.Type("url")

		default:
			tag.Type("text")
		}

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

	default:
		tag.Type("text")
	}

	if autocomplete := element.Options.GetString("autocomplete"); autocomplete != "" {
		tag.Attr("autocomplete", autocomplete)

		if autocomplete == "off" {
			tag.Attr("data-1p-ignore", "true")
		}
	}

	tag.Value(valueString)
	tag.TabIndex("0")
	tag.Close()

	if len(lookupCodes) > 0 {
		b.Container("datalist").ID("datalist-" + element.Path)
		for _, lookupCode := range lookupCodes {
			b.Empty("option").Value(lookupCode.Value).Close() // Datalist options do not have an innerHTML
		}
		b.Close()
	}

	b.CloseAll()
	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget Text) ShowLabels() bool {
	return true
}

func (widget Text) Encoding(_ *form.Element) string {
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
