package widget

import (
	"time"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/first"
	"github.com/benpate/rosetta/schema"
)

type DatePicker struct{}

func (widget DatePicker) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)

	// TODO: LOW: Apply formatting options?
	b.Div().Class("layout-value", e.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

func (widget DatePicker) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	valueString := widget.getValue(e, &f.Schema, value)
	eID := first.String(e.ID, "datepicker-"+e.Path)

	// Start building a new tag
	tag := b.Input("date", e.Path).
		ID(eID).
		Value(valueString).
		TabIndex("0")

	if focus, ok := e.Options.GetBoolOK("focus"); ok && focus {
		tag.Attr("autofocus", "true")
	}

	b.CloseAll()
	return nil
}

func (widget DatePicker) getValue(e *form.Element, s *schema.Schema, value any) string {

	valueString := e.GetString(value, s)

	if result, ok := convert.TimeOk(valueString, time.Time{}); ok {
		return result.Format("2006-01-02")
	}

	return ""
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget DatePicker) ShowLabels() bool {
	return true
}

func (widget DatePicker) Encoding(_ *form.Element) string {
	return ""
}
