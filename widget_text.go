package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("text", WidgetText{})
}

type WidgetText struct{}

func (WidgetText) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := element.GetString(value, s)

	// TODO: LOW: Apply formatting options?
	b.Div().Class("layout-value").InnerText(valueString).Close()
	return nil
}

func (WidgetText) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetText{}.View(element, s, lookupProvider, value, b)
	}

	// find the path and schema to use
	schemaElement := element.getElement(s)
	valueString := element.GetString(value, s)

	elementID := element.ID

	if elementID == "" {
		elementID = "text-" + element.Path
	}

	// Start building a new tag
	tag := b.Input("", element.Path).
		ID(elementID).
		Value(valueString)

	if focus, ok := element.Options.GetBoolOK("focus"); ok && focus {
		tag.Attr("autofocus", "true")
	}

	// Enumeration Options
	lookupCodes, _ := GetLookupCodes(element, schemaElement, lookupProvider)
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
		case "email":
			tag.Type("email")
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
	}
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

func (WidgetText) ShowLabels() bool {
	return true
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
