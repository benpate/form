package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
)

type Password struct{}

func (widget Password) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	b.Div().Class("layout-value", e.Options.GetString("class")).InnerText("********").Close()
	return nil
}

func (widget Password) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	if e.ID == "" {
		e.ID = e.Path + "." + e.Type
	}

	// Start building a new tag
	tag := b.Input("password", e.Path).
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

	// Custom CSS style
	if style := e.Options.GetString("style"); style != "" {
		tag.Attr("style", style)
	}

	// Password rules may not have a schema attached, so use options instead.
	if minlength := e.Options.GetInt("minlength"); minlength > 0 {
		tag.Attr("minlength", convert.String(minlength))
	}

	if maxlength := e.Options.GetInt("maxlength"); maxlength > 0 {
		tag.Attr("maxlength", convert.String(maxlength))
	}

	if pattern := e.Options.GetString("pattern"); pattern != "" {
		tag.Attr("pattern", pattern)
	}

	if required := e.Options.GetBool("required"); required {
		tag.Attr("required", "true")
	}

	if autocomplete := e.Options.GetString("autocomplete"); autocomplete != "" {
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
