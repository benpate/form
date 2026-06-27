package widget

import (
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/compare"
)

// CheckButton renders a fancy checkbox widget that looks like a button
type CheckButton struct{}

// View generates the read-only HTML for this check-button (which is nothing).
func (widget CheckButton) View(_ *form.Form, _ *form.Element, _ form.LookupProvider, _ any, _ *html.Builder) error {
	return nil
}

// Edit generates the editable HTML for this check-button.
func (widget CheckButton) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	// Collect values to use in the Widget
	selectedValues, err := f.Schema.Get(value, e.Path)

	if err != nil {
		derp.Report(derp.Wrap(err, "form.checkbutton.Edit", "Error getting value for CheckButton", e.Path, value))
	}

	elementValue := e.Options.GetString("value")
	id := "checkbutton-" + strings.ReplaceAll(e.Path, ".", "-") + "-" + elementValue

	// Build the widget HTML
	b.Label(id).ID("label-" + id).Class("checkbutton")

	if icon := e.Options.GetString("icon"); icon != "" {
		b.I().Class("margin-horizontal", "bi", "bi-"+icon).Style("font-size:32px;").Close()
	}

	b.Div().Class("flex-column")
	b.Div().Class("bold").InnerText(e.Label).Close()
	b.Div().Class("text-sm", "text-gray").InnerText(e.Description).Close()

	checkbox := b.Input("checkbox", e.Path)
	checkbox.
		ID(id).
		Value(elementValue).
		Class(e.Options.GetString("class")).
		Script(e.Options.GetString("script")).
		Aria("label", e.Label).
		Aria("description", e.Description).
		TabIndex("0")

	if compare.Contains(selectedValues, elementValue) {
		checkbox.Attr("checked", "true")
	}

	if disabled := e.Options.GetBool("disabled"); disabled {
		checkbox.Attr("disabled", "true")
	}

	b.CloseAll()
	return nil
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For CheckButton widgets, the label is drawn inside the button, so this always returns FALSE.
func (widget CheckButton) ShowLabels() bool {
	return false
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget CheckButton) ShowDescriptions() string {
	return "NONE"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For CheckButton widgets, there is no special encoding,
// so this always returns an empty string.
func (widget CheckButton) Encoding(_ *form.Element) string {
	return ""
}
