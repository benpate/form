package widget

import (
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

// Upload is a widget that draws a single-file upload zone: the saved file as a tile, or a plus that picks one.
type Upload struct{}

// View generates the read-only HTML for this uploaded file's value.
func (widget Upload) View(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := e.GetString(value, &f.Schema)

	if valueString == "" {
		valueString = "N/A"
	}

	b.Div().Class("layout-value", e.Options.GetString("class", value)).InnerText(valueString).Close()
	return nil
}

// Edit generates the editable HTML for this upload zone: the saved file's tile, the plus, the file input, and a status line.
func (widget Upload) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	const location = "widget.Upload.Edit"

	elementID := e.ID

	if elementID == "" {
		elementID = e.Path + ".upload"
	}

	valueString := e.GetString(value, &f.Schema)
	accept := e.Options.GetString("accept", value)

	// The zone is marked for the host's styles and behavior, which find it by this class and this attribute
	classes := []string{"upload-single"}

	if class := e.Options.GetString("class", value); class != "" {
		classes = append(classes, class)
	}

	b.Div().Class(classes...).Data("upload-single", "true")

	// Draw the saved file, when there is one
	if err := widget.tile(e, &f.Schema, value, b.SubTree()); err != nil {
		return derp.Wrap(err, location, "Unable to draw the saved file")
	}

	// RULE: The plus is type="button" because a bare button inside a form submits it.
	// It gives way to the saved file's tile, which picks a replacement itself.
	b.Button().
		Type("button").
		Class("upload-multiple-add").
		Attr("title", "Choose a file").
		Attr("hidden", iif(valueString == "", "", "hidden")).
		Aria("label", "Choose a file")
	b.I("bi", "bi-plus-lg").Close()
	b.Close()

	// RULE: The input is hidden by the host's styles, never by an attribute, so a page with no styles still has a working picker
	multiple := iif(e.Options.GetBool("multiple", value), "multiple", "")
	b.Input("file", e.Path).ID(elementID).
		Attr("accept", accept).
		Attr("multiple", multiple).
		Aria("label", e.Label).
		Aria("description", e.Description).
		Close()

	b.Div().Class("upload-multiple-status").Role("status").Close()
	b.Close()

	// A zone is born
	return nil
}

// tile draws the already-uploaded file as a tile: an image preview, or an icon and a filename
// for anything else, with the X that deletes it when the "delete" option names a URL.
func (widget Upload) tile(e *form.Element, s *schema.Schema, value any, b *html.Builder) error {

	const location = "widget.Upload.tile"

	// Get the URL for the uploaded file
	valueString := e.GetString(value, s)

	if valueString == "" {
		return nil
	}

	// The delete URL is usually an option template naming the object being edited
	deleteLink, err := e.Options.Evaluate("delete", value)

	if err != nil {
		return derp.Wrap(err, location, "Unable to render delete URL")
	}

	filename := widget.filename(e, s, value, valueString)

	// The tile is a button in all but name: a click, Enter, or Space on it picks a replacement
	b.Div().
		Class("upload-multiple-tile", "hover-trigger").
		Data("upload-single-saved", "true").
		Attr("tabindex", "0").
		Aria("label", "Replace "+filename)

	// RULE: The filename and the URL are end-user data, so they are written through Attr and InnerText, which escape
	if acceptType, _, _ := strings.Cut(e.Options.GetString("accept", value), "/"); acceptType == "image" {
		b.Img(valueString).Class("upload-multiple-preview").Attr("alt", filename).Attr("draggable", "false").Close()
	} else {
		b.I("bi", "bi-file-earmark", "upload-multiple-icon").Close()
		b.Span().Class("upload-multiple-filename").InnerText(filename).Close()
	}

	// The saved URL posts with the form, so a Save with no new file keeps it
	b.Input("hidden", e.Path).Value(valueString).Close()

	if deleteLink != "" {
		b.Button().
			Type("button").
			Class("upload-multiple-remove", "hover-show").
			Attr("hx-post", deleteLink).
			Attr("hx-confirm", "Delete this file?").
			Attr("hx-swap", "none").
			Attr("hx-push-url", "false").
			Aria("label", "Delete "+filename)
		b.I("bi", "bi-x-lg").Close()
		b.Close()
	}

	b.Close()

	// Bob's your uncle
	return nil
}

// filename returns the name to show for the saved file: the value at the "filename" option's
// path when it has one, or else the last segment of the file's URL.
func (widget Upload) filename(e *form.Element, s *schema.Schema, value any, url string) string {

	// The "filename" option names the path where the host keeps the file's original name
	if filenamePath := e.Options.GetString("filename", value); filenamePath != "" {
		if original, err := s.Get(value, filenamePath); err == nil {
			if result := form.FormatValue(original); result != "" {
				return result
			}
		}
	}

	// Otherwise, use the URL's last segment, without its query or fragment
	result, _, _ := strings.Cut(url, "?")
	result, _, _ = strings.Cut(result, "#")
	result = strings.TrimRight(result, "/")

	if index := strings.LastIndex(result, "/"); index >= 0 {
		result = result[index+1:]
	}

	if result == "" {
		return "file"
	}

	return result
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
