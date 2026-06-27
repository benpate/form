package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
)

// Label is a widget that displays static label and description text without an input.
type Label struct{}

// View generates the read-only HTML for this label.
func (widget Label) View(_ *form.Form, e *form.Element, _ form.LookupProvider, _ any, b *html.Builder) error {
	if e.Label != "" {
		b.Div().InnerText(e.Label).Close()
	}

	if e.Description != "" {
		b.Div().InnerHTML(e.Description).Close()
	}

	return nil
}

// Edit generates the HTML for this label, which is identical to its read-only view.
func (widget Label) Edit(_ *form.Form, e *form.Element, _ form.LookupProvider, _ any, b *html.Builder) error {
	if e.Label != "" {
		b.Div().InnerText(e.Label).Close()
	}

	if e.Description != "" {
		b.Div().InnerHTML(e.Description).Close()
	}

	return nil
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For Label widgets, the text is the content itself, so this always returns FALSE.
func (widget Label) ShowLabels() bool {
	return false
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget Label) ShowDescriptions() string {
	return "NONE"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// For Label widgets, there is no special encoding,
// so this always returns an empty string.
func (widget Label) Encoding(_ *form.Element) string {
	return ""
}
