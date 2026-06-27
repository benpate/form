package widget

import (
	"time"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/first"
	"github.com/benpate/rosetta/schema"
)

// DatePicker is a widget that creates a date input field.
type DatePicker struct{}

// View generates the read-only HTML for this date value.
func (widget DatePicker) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)

	// TODO: LOW: Apply formatting options?
	b.Div().Class("layout-value", e.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

// Edit generates the editable HTML for this date input field.
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

	if defaultValue := e.Options.GetString("default"); defaultValue != "" {

		if defaultValue == "now" {
			tag.Script(`on load if my value is "" then make a Date set my valueAsDate to it`)
		}
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
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For DatePicker widgets, labels are shown, so this always returns TRUE.
func (widget DatePicker) ShowLabels() bool {
	return true
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget DatePicker) ShowDescriptions() string {
	return "BOTTOM"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For DatePicker widgets, there is no special encoding,
// so this always returns an empty string.
func (widget DatePicker) Encoding(_ *form.Element) string {
	return ""
}
