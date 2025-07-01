package widget

import (
	"strings"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

type Text struct{}

func (widget Text) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)

	// TODO: LOW: Apply formatting options?
	b.Div().Class("layout-value", e.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

func (widget Text) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	if e.ReadOnly {
		return Text{}.View(f, e, provider, value, b)
	}

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)
	valueString := e.GetString(value, &f.Schema)

	if e.ID == "" {
		e.ID = e.Path + "." + e.Type
	}

	scripts := make([]string, 0)

	if validator := e.Options.GetString("validator"); validator != "" {
		b.Div().Class("badge-container").EndBracket()
		scripts = append(scripts, `install validator(url:'`+validator+`')`)
	}

	// Start building a new tag
	tag := b.Input("", e.Path).
		ID(e.ID).
		Aria("label", e.Label).
		Aria("description", e.Description).
		TabIndex("0")

	if focus, ok := e.Options.GetBoolOK("focus"); ok && focus {
		tag.Attr("autofocus", "true")
	}

	if placeholder := e.Options.GetString("placeholder"); placeholder != "" {
		tag.Attr("placeholder", placeholder)
	}

	// Enumeration Options
	lookupCodes, _ := form.GetLookupCodes(e, schemaElement, provider)
	if len(lookupCodes) > 0 {
		tag.Attr("list", "datalist-"+e.Path)
	}

	// Custom CSS style
	if style := e.Options.GetString("style"); style != "" {
		tag.Attr("style", style)
	}

	// Add attributes that depend on what KIND of input we have.
	switch s := schemaElement.(type) {

	case schema.Integer:

		tag.Type("number")

		tag.Attr("step", convert.String(convert.IntDefault(e.Options["step"], 1)))

		if s.Minimum.IsPresent() {
			tag.Attr("min", s.Minimum.String())
		}

		if s.Maximum.IsPresent() {
			tag.Attr("max", s.Maximum.String())
		}

		if s.Required || e.Options.GetBool("required") {
			tag.Attr("required", "true")
		}

		if s.RequiredIf != "" {
			scripts = append(scripts, "install requiredIf(condition:'"+s.RequiredIf+"')")
		} else if requiredIf := e.Options.GetString("required-if"); requiredIf != "" {
			scripts = append(scripts, "install requiredIf(condition:'"+requiredIf+"')")
		}

	case schema.Number:

		tag.Type("number")

		tag.Attr("step", convert.String(convert.FloatDefault(e.Options["step"], 0.01)))

		if s.Minimum.IsPresent() {
			tag.Attr("min", s.Minimum.String())
		}

		if s.Maximum.IsPresent() {
			tag.Attr("max", s.Maximum.String())
		}

		if s.Required || e.Options.GetBool("required") {
			tag.Attr("required", "true")
		}

		if s.RequiredIf != "" {
			scripts = append(scripts, "install requiredIf(condition:'"+s.RequiredIf+"')")
		} else if requiredIf := e.Options.GetString("required-if"); requiredIf != "" {
			scripts = append(scripts, "install requiredIf(condition:'"+requiredIf+"')")
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
		} else if pattern := e.Options.GetString("pattern"); pattern != "" {
			tag.Attr("pattern", pattern)
		}

		if s.Required || e.Options.GetBool("required") {
			tag.Attr("required", "true")
		}

		if s.RequiredIf != "" {
			scripts = append(scripts, "install requiredIf(condition:'"+s.RequiredIf+"')")
		} else if requiredIf := e.Options.GetString("required-if"); requiredIf != "" {
			scripts = append(scripts, "install requiredIf(condition:'"+requiredIf+"')")
		}

	default:
		tag.Type("text")
	}

	if autocomplete := e.Options.GetString("autocomplete"); autocomplete != "" {
		tag.Attr("autocomplete", autocomplete)

		if autocomplete == "off" {
			tag.Attr("data-1p-ignore", "true")
		}
	}

	if autocorrect := e.Options.GetString("autocorrect"); autocorrect != "" {
		tag.Attr("autocorrect", autocorrect)
	}

	if spellcheck := e.Options.GetString("spellcheck"); spellcheck != "" {
		tag.Attr("spellcheck", spellcheck)
	}

	if len(scripts) > 0 {
		tag.Script(strings.Join(scripts, " "))
	}

	tag.Value(valueString)
	tag.Close()

	if len(lookupCodes) > 0 {
		b.Container("datalist").ID("datalist-" + e.Path)
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
