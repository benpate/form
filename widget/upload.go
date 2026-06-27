package widget

import (
	"strings"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

// Upload is a widget that creates a file-upload input with a preview of the current file.
type Upload struct{}

// View generates the read-only HTML for this uploaded file's value.
func (widget Upload) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)

	if valueString == "" {
		valueString = "N/A"
	}

	b.Div().Class("layout-value", e.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

// Edit generates the editable HTML for this file-upload input.
func (widget Upload) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	elementID := e.ID

	if elementID == "" {
		elementID = e.Path + ".upload"
	}

	widget.preview(e, &f.Schema, value, b.SubTree())

	multiple := iif(e.Options.GetBool("multiple"), "multiple", "")
	b.Input("file", e.Path).ID(elementID).
		Attr("accept", e.Options.GetString("accept")).
		Attr("multiple", multiple).
		Aria("label", e.Label).
		Aria("description", e.Description).
		TabIndex("0").
		Close()

	return nil
}

func (widget Upload) preview(e *form.Element, s *schema.Schema, value any, b *html.Builder) {

	// Get the URL for the uploaded file
	valueString := e.GetString(value, s)

	if valueString == "" {
		return
	}

	// Different file types are displayed differently
	accept := e.Options.GetString("accept")

	b.Div().Class("pos-relative", "width-128").Style("border:solid 1px black")

	switch acceptType, _, _ := strings.Cut(accept, "/"); acceptType {

	// Image preview (128px square)
	case "image":
		b.Img(valueString).Style("display:block", "width:128px", "height:128px", "object-fit:cover").Close()

	// Audio previoew (with controls)
	case "audio":
		b.Audio().Attr("controls", "true")
		b.Source().Src(valueString).Close()
		b.Close()

	// All other files are displayed as a link
	default:
		b.A(valueString).InnerText(valueString).Close()
	}

	b.Input("hidden", e.Path).Value(e.GetString(value, s)).Close()
	if deleteLink := e.Options.GetString("delete"); deleteLink != "" {
		b.Span().
			Class("pos-absolute-top-right text-xs button").
			Attr("hx-post", deleteLink).
			Attr("hx-confirm", "Delete this file?").
			Attr("script", "on htmx:afterRequest remove my parentNode").
			Aria("label", "Delete").
			InnerText("X").
			Close()
	}
	b.Close()
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For Upload widgets, labels are shown, so this always returns TRUE.
func (widget Upload) ShowLabels() bool {
	return true
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget Upload) ShowDescriptions() string {
	return "BOTTOM"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// Upload widgets submit files, so this always returns "multipart/form-data".
func (widget Upload) Encoding(_ *form.Element) string {
	return "multipart/form-data"
}
