package widget

import (
	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
)

// HTMLRemote is a widget that loads its content from a remote URL, templated against the
// form's value, as soon as it is on the page. It draws nothing of its own and binds nothing
// to the object; whatever the URL answers with is the field.
type HTMLRemote struct{}

// View generates the read-only HTML for this remote-loading element,
// which is identical to its editable form.
func (widget HTMLRemote) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	return widget.Edit(f, e, nil, value, b)
}

// Edit generates the editable HTML for this remote-loading element: one empty div, with the
// element's ID so a label can name it, that htmx fills from the URL when it loads.
func (widget HTMLRemote) Edit(_ *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	const location = "widget.HTMLRemote.Edit"

	// Render the Remote URL, which may be an option template
	remoteURL, err := e.Options.Evaluate("url", value)

	if err != nil {
		return derp.Wrap(err, location, "Unable to render remote URL")
	}

	b.Div().
		ID(e.ID).
		Attr("hx-get", remoteURL).
		Attr("hx-swap", "innerHTML").
		Attr("hx-trigger", "load").
		Attr("hx-target", "this").
		Attr("hx-push-url", "false").
		Close()

	return nil
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It indicates whether or not this widget wants labels to be shown.
// For HTMLRemote widgets, labels are shown, as for any other field.
func (widget HTMLRemote) ShowLabels() bool {
	return true
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget HTMLRemote) ShowDescriptions() string {
	return "BOTTOM"
}

// Encoding is a part of the Widget interface
// It returns the encoding type for this widget
// For HTMLRemote widgets, it always returns empty string
func (widget HTMLRemote) Encoding(_ *form.Element) string {
	return ""
}
