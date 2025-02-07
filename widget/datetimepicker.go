package widget

import (
	"time"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/first"
	"github.com/benpate/rosetta/schema"
)

type DateTimePicker struct{}

func (widget DateTimePicker) View(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := element.GetString(value, s)

	// TODO: LOW: Apply formatting options?
	b.Div().Class("layout-value", element.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

func (widget DateTimePicker) Edit(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return DateTimePicker{}.View(element, s, nil, value, b)
	}

	valueString := widget.getValue(element, s, value)
	elementID := first.String(element.ID, "datetimepicker-"+element.Path)

	// Start building a new tag
	tag := b.Input("datetime-local", element.Path).
		ID(elementID).
		Value(valueString).
		TabIndex("0")

	if focus, ok := element.Options.GetBoolOK("focus"); ok && focus {
		tag.Attr("autofocus", "true")
	}

	b.CloseAll()
	return nil
}

func (widget DateTimePicker) getValue(element *form.Element, s *schema.Schema, value any) string {

	valueString := element.GetString(value, s)

	if result, ok := convert.TimeOk(valueString, time.Time{}); ok {
		return result.Format("2006-01-02T15:04")
	}

	return ""
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget DateTimePicker) ShowLabels() bool {
	return true
}

func (widget DateTimePicker) Encoding(_ *form.Element) string {
	return ""
}
