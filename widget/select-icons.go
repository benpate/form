package widget

import (
	"strconv"

	"github.com/benpate/form"
	"github.com/benpate/form/groupie"
	"github.com/benpate/html"
)

// SelectIcons is a widget that chooses one icon from a popover grid of tiles.
type SelectIcons struct{}

// View generates the read-only HTML for this widget, showing the selected icon and its Label.
func (widget SelectIcons) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)
	valueString := e.GetString(value, &f.Schema)
	lookupCodes, _ := form.GetLookupCodes(e, schemaElement, provider)
	selected := selectedCode(lookupCodes, valueString)

	// RULE: "layout-value" alone. view_escaping_test.go asserts this exact opening
	// tag across every lookup widget, which is what proves no text leaks into it.
	b.Div().Class("layout-value")

	if name := iconName(selected); name != "" {
		b.I("bi", "bi-"+name).Close()
	}

	// InnerText escapes: a Label may hold whatever an end-user typed
	b.Span().InnerText(selected.Label).Close()

	b.CloseAll()
	return nil
}

// Edit generates the editable HTML for this widget: a trigger button plus a popover of radio tiles.
func (widget SelectIcons) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	schemaElement := e.GetSchema(&f.Schema)
	valueString := e.GetString(value, &f.Schema)
	lookupCodes, _ := form.GetLookupCodes(e, schemaElement, provider)

	elementID := getElementID(e)
	popoverID := elementID + "-popover"

	// RULE: When the field is required, an unmatched value falls back to the first
	// option, which is what a <select> with no blank <option> already does.
	required := isRequired(e, schemaElement)
	selectedIndex := indexOfCode(lookupCodes, valueString)

	if (selectedIndex < 0) && required && (len(lookupCodes) > 0) {
		selectedIndex = 0
	}

	b.Div().ID(elementID).Class("select-icons").Script("install selectIcons")

	widget.drawButton(b, e, elementID, popoverID, lookupCodes, selectedIndex)

	// RULE: "auto" and not an empty value, because Attr skips empty values
	// entirely -- which would drop the attribute and leave the panel visible.
	b.Div().
		ID(popoverID).
		Class("select-icons-popover").
		Style("position-anchor:--"+elementID).
		Attr("popover", "auto")
	b.Div().Class("select-icons-grid").Role("radiogroup").Aria("label", e.Label)

	// RULE: A blank option is offered only when the field is not required,
	// matching the blank <option> that the Select widget emits.
	if !required {
		widget.drawTile(b, e, elementID, form.LookupCode{Label: "None", Icon: "slash-circle"}, -1, selectedIndex < 0)
	}

	group := groupie.New()

	for index, lookupCode := range lookupCodes {

		if group.Header(lookupCode.Group) && (lookupCode.Group != "") {
			b.Div().Class("select-icons-header").InnerText(lookupCode.Group).Close()
		}

		widget.drawTile(b, e, elementID, lookupCode, index, index == selectedIndex)
	}

	b.CloseAll()
	return nil
}

// drawButton writes the trigger button that opens the popover.
func (widget SelectIcons) drawButton(b *html.Builder, e *form.Element, elementID string, popoverID string, lookupCodes []form.LookupCode, selectedIndex int) {

	selected := form.LookupCode{Label: "None"}

	if (selectedIndex >= 0) && (selectedIndex < len(lookupCodes)) {
		selected = lookupCodes[selectedIndex]
	}

	// RULE: Inside a <form> a bare <button> submits. popovertarget changes what
	// activation does; it does not stop a submit button from submitting.
	// RULE: The anchor name is per-widget. A name shared in the stylesheet would
	// collide wherever a form draws two icon pickers, as user-inbox does.
	//
	// The "input" class is what gives the trigger the host theme's own input
	// padding, border, and radius, so it tracks them instead of copying them.
	b.Button().
		ID(elementID+"-button").
		Type("button").
		Class("input", "select-icons-button").
		Style("anchor-name:--"+elementID).
		Attr("popovertarget", popoverID).
		Aria("label", e.Label).
		Aria("description", e.Description)

	widget.drawIcon(b, selected, "select-icons-button-icon")

	// InnerText escapes: a Label may hold whatever an end-user typed
	b.Span().Class("select-icons-button-label").InnerText(selected.Label).Close()
	b.I("bi", "bi-chevron-down").Close()

	b.Close()
}

// drawTile writes one <label> wrapping the radio input for a single LookupCode.
func (widget SelectIcons) drawTile(b *html.Builder, e *form.Element, elementID string, lookupCode form.LookupCode, index int, checked bool) {

	// RULE: Tiles are keyed by index, never by Value. A Value may be an emoji or
	// any other string, and none of them are safe to paste into a DOM id.
	tileID := elementID + "-" + strconv.Itoa(index+1)

	// RULE: data-icon carries the resolved name so the behavior can rebuild the
	// button face without re-deriving it from a Value that may not be an icon.
	b.Label(tileID).
		Class("select-icons-tile").
		Attr("title", lookupCode.Label).
		Data("icon", iconName(lookupCode))

	// RULE: The tile shows no text, so aria-label is the only accessible name it
	// has. The title attribute is the sighted equivalent, on hover.
	radio := b.Input("radio", e.Path).
		ID(tileID).
		Value(lookupCode.Value).
		Aria("label", lookupCode.Label)

	if checked {
		radio.Attr("checked", "true")
	}

	radio.Close()

	widget.drawIcon(b, lookupCode, "select-icons-tile-icon")

	b.Close()
}

// drawIcon writes the <i> tag for a LookupCode, or an empty placeholder of the
// same size when the code names no usable icon.
func (widget SelectIcons) drawIcon(b *html.Builder, lookupCode form.LookupCode, class string) {

	// RULE: Every class goes in the one I() call. A second .Class() writes a
	// second class attribute, and the browser keeps only the first.
	if name := iconName(lookupCode); name != "" {
		b.I("bi", "bi-"+name, class).Close()
		return
	}

	b.Span().Class(class, "select-icons-blank").Close()
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For SelectIcons widgets, labels are shown, so this always returns TRUE.
func (widget SelectIcons) ShowLabels() bool {
	return true
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget SelectIcons) ShowDescriptions() string {
	return "BOTTOM"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For SelectIcons widgets, there is no special encoding,
// so this always returns an empty string.
func (widget SelectIcons) Encoding(_ *form.Element) string {
	return ""
}

// indexOfCode returns the position of the first LookupCode whose Value matches
// the one provided, or -1 when nothing matches.
func indexOfCode(lookupCodes []form.LookupCode, value string) int {

	for index, lookupCode := range lookupCodes {
		if lookupCode.Value == value {
			return index
		}
	}

	return -1
}
