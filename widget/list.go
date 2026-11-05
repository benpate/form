package widget

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/first"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/sliceof"
)

// List is a widget that edits an ordered list of strings, one text input per row.
type List struct{}

// Edit generates the editable HTML for this list: a row per value, a blank row to copy, and an add button.
func (widget List) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	const location = "widget.List.Edit"

	// RULE: A read-only list cannot be edited, so it draws its View instead
	if e.ReadOnly {
		return widget.View(f, e, provider, value, b)
	}

	// RULE: The path must hold an Array of Strings, so a mismatch fails here instead of never saving
	arraySchema, itemSchema, err := listSchema(f, e)

	if err != nil {
		return derp.Wrap(err, location, "Invalid schema for list", e.Path)
	}

	// An empty list draws one blank row, so people can type at once
	values := e.GetSliceOfString(value, &f.Schema)

	if len(values) == 0 {
		values = []string{""}
	}

	// Name the wrapper, and the hint that every row's input points to
	elementID := getElementID(e)
	hintID := elementID + "-hint"

	// RULE: The X is disabled at the Array's minimum, as the table's delete is
	canDelete := len(values) > arraySchema.MinLength

	// Wrapper, with the limits the behavior keeps up to date
	wrapper := b.Div().ID(elementID).Class("list-widget").Script("install ListWidget")

	if arraySchema.MinLength > 0 {
		wrapper.Data("minimum", strconv.Itoa(arraySchema.MinLength))
	}

	if arraySchema.MaxLength > 0 {
		wrapper.Data("maximum", strconv.Itoa(arraySchema.MaxLength))
	}

	if isRequired(e, arraySchema, value) {
		wrapper.Data("required", "true")
	}

	wrapper.EndBracket()

	// One row per value
	b.Table().Class("grid", "list")
	b.TBody()

	for _, item := range values {
		widget.drawRow(b, e, itemSchema, item, hintID, canDelete)
	}

	b.Close() // TBODY
	b.Close() // TABLE

	// A blank row for the behavior to copy when it adds one
	b.Container("template").Class("list-template")
	widget.drawRow(b, e, itemSchema, "", hintID, true)
	b.Close() // TEMPLATE

	// RULE: The add button goes at the Array's maximum, as the table's does
	addButton := b.Button().Type("button").Class("link", "list-add")

	if (arraySchema.MaxLength > 0) && (len(values) >= arraySchema.MaxLength) {
		addButton.Attr("hidden", "true")
	}

	addButton.EndBracket()
	b.I("bi", "bi-plus-lg").Aria("hidden", "true").Close()
	b.Span().InnerText(" " + first.String(e.Options.GetString("add-label", value), "Add a Row")).Close()
	b.Close() // BUTTON

	// The keys, for screen readers, and the polite live region that announces each move
	b.Div().ID(hintID).Class("list-hint", "hide-visual").InnerText("Enter adds a row below. Alt+Up and Alt+Down move this row, and Alt+Home and Alt+End move it to the top or bottom.").Close()
	b.Div().Class("list-status", "hide-visual").Aria("live", "polite").Close()

	b.CloseAll()

	// Here's a list of things I need to do.  1. Make a list.
	return nil
}

// View generates the read-only HTML for this list, one row of text per value.
func (widget List) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	const location = "widget.List.View"

	// RULE: The path must hold an Array of Strings, as it must for Edit
	if _, _, err := listSchema(f, e); err != nil {
		return derp.Wrap(err, location, "Invalid schema for list", e.Path)
	}

	// An empty list draws the same empty value as every other widget
	values := e.GetSliceOfString(value, &f.Schema)

	if len(values) == 0 {
		drawValue(b, "")
		return nil
	}

	// Draw each value as one row of text
	b.Table().Class("grid", "list")
	b.TBody()

	for _, item := range values {
		b.TR().Class("grid-row")
		b.TD().Class("grid-cell").InnerText(item).Close()
		b.Close() // TR
	}

	b.CloseAll()

	// Nothing to see here
	return nil
}

// SetURLValue is a part of the URLValueSetter interface.
// It writes the posted rows, trimmed and without blanks or repeats, as the list's new value.
func (widget List) SetURLValue(f *form.Form, e *form.Element, object any, values url.Values) error {

	const location = "widget.List.SetURLValue"

	// RULE: Never write a path that is not an Array of Strings
	if _, _, err := listSchema(f, e); err != nil {
		return derp.Wrap(err, location, "Invalid schema for list", e.Path)
	}

	// Write the cleaned rows through a *sliceof.String, as form.schemaSafeValue writes an Array path
	result := cleanList(values[e.Path])

	if err := f.Schema.Set(object, e.Path, &result); err != nil {
		return derp.Wrap(err, location, "Unable to set list", e.Path)
	}

	// Check it twice
	return nil
}

// listSchema returns the Array at this element's path and the String that each of its items must be,
// or an error when the path holds anything else.
func listSchema(f *form.Form, e *form.Element) (schema.Array, schema.String, error) {

	const location = "widget.listSchema"

	arraySchema, isArray := e.GetSchema(&f.Schema).(schema.Array)

	if !isArray {
		return schema.Array{}, schema.String{}, derp.Internal(location, "List path must be an Array", e.Path)
	}

	itemSchema, isString := arraySchema.Items.(schema.String)

	if !isString {
		return schema.Array{}, schema.String{}, derp.Internal(location, "List items must be Strings", e.Path)
	}

	return arraySchema, itemSchema, nil
}

// drawRow writes one editable row: the drag grip, the text input, and the X that removes it.
func (widget List) drawRow(b *html.Builder, e *form.Element, itemSchema schema.String, value string, hintID string, canDelete bool) {

	// Each row is a table row, which is what the ListWidget behavior and Sortable move
	b.TR().Class("grid-row")

	// The grip is a pointer target only; the keys move a row from its input
	b.TD().Class("grid-cell", "list-handle").Aria("hidden", "true")
	b.I("bi", "bi-grip-vertical").Close()
	b.Close() // TD

	// RULE: No "required" on a row, because a blank row is allowed and dropped at Save.
	// Browsers offer autofill for repeated names, which would put old values in every new row.
	b.TD().Class("grid-cell")
	input := b.Input("text", e.Path).
		Value(value).
		Aria("label", e.Label).
		Aria("describedby", hintID).
		Attr("autocomplete", "off")

	// The item String's limits go on every input, so the browser checks each row
	if itemSchema.MinLength > 0 {
		input.Attr("minlength", strconv.Itoa(itemSchema.MinLength))
	}

	if itemSchema.MaxLength > 0 {
		input.Attr("maxlength", strconv.Itoa(itemSchema.MaxLength))
	}

	input.Attr("pattern", itemSchema.Pattern)

	b.Close() // INPUT
	b.Close() // TD

	// RULE: Inside a form a bare <button> submits, so the X is always type="button"
	b.TD().Class("grid-cell", "grid-controls")
	remove := b.Button().Type("button").Class("list-remove").Aria("label", "Remove this row")

	if !canDelete {
		remove.Attr("disabled", "true")
	}

	remove.EndBracket()
	b.I("bi", "bi-x-lg").Aria("hidden", "true").Close()
	b.Close() // BUTTON
	b.Close() // TD

	b.Close() // TR
}

// cleanList trims each value, drops blanks, and drops any value that repeats an earlier one ignoring case.
func cleanList(values []string) sliceof.String {

	result := make(sliceof.String, 0, len(values))
	seen := make(map[string]struct{}, len(values))

	for _, value := range values {

		value = strings.TrimSpace(value)

		if value == "" {
			continue
		}

		// RULE: The first copy wins, so the order people chose is kept.
		// The behavior folds case with toLowerCase(), which differs only for a few characters.
		key := strings.ToLower(value)

		if _, isRepeat := seen[key]; isRepeat {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, value)
	}

	return result
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE, because a list is drawn with its label.
func (widget List) ShowLabels() bool {
	return true
}

// ShowDescriptions is a part of the Widget interface.
// It returns "BOTTOM", so the description is drawn below the list.
func (widget List) ShowDescriptions() string {
	return "BOTTOM"
}

// Encoding is a part of the Widget interface.
// It returns an empty string, because a list needs no special encoding.
func (widget List) Encoding(_ *form.Element) string {
	return ""
}
