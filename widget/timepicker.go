package widget

import (
	"time"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/first"
	"github.com/benpate/rosetta/schema"
)

type TimePicker struct{}

func (widget TimePicker) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)

	// TODO: LOW: Apply formatting options?
	b.Div().Class("layout-value", e.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

func (widget TimePicker) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	valueString := widget.getValue(e, &f.Schema, value)
	elementID := first.String(e.ID, "timepicker-"+e.Path)

	// Start building a new tag
	tag := b.Input("time", e.Path).
		ID(elementID).
		Value(valueString).
		TabIndex("0")

	if focus, ok := e.Options.GetBoolOK("focus"); ok && focus {
		tag.Attr("autofocus", "true")
	}

	if defaultValue := e.Options.GetString("default"); defaultValue != "" {

		if defaultValue == "now" {
			tag.Script(`on load if my value is "" then make a Date set my valueAsDate to it`)
		}
	}

	b.CloseAll()
	return nil
}

func (widget TimePicker) getValue(e *form.Element, s *schema.Schema, value any) string {

	valueString := e.GetString(value, s)

	if result, ok := convert.TimeOk(valueString, time.Time{}); ok {
		return result.Format("15:04")
	}

	return valueString
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget TimePicker) ShowLabels() bool {
	return true
}

func (widget TimePicker) Encoding(_ *form.Element) string {
	return ""
}
