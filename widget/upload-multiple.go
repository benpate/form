package widget

import (
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
)

// UploadMultiple is a widget that takes several files in one field, one at a time: a tile row
// that loads the saved files from the host's own route, a plus tile that opens a single-file
// picker, and a hidden file input. It binds nothing to the object; the host's step stores,
// deletes, and sorts the files.
type UploadMultiple struct{}

// View generates the read-only HTML for this widget: the tile row, loaded from the host's route.
func (widget UploadMultiple) View(_ *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	const location = "widget.UploadMultiple.View"

	remoteURL, err := widget.remoteURL(e, value)

	if err != nil {
		return derp.Wrap(err, location, "Unable to render url")
	}

	// RULE: "layout-value" alone. view_escaping_test.go asserts this exact opening
	// tag across the read-only widgets, which is what proves no text leaks into it.
	b.Div().Class("layout-value")
	widget.drawTileRow(b, widget.elementID(e, value), remoteURL)
	b.CloseAll()

	return nil
}

// Edit generates the editable HTML for this widget: the zone, the tile row, the plus tile
// that opens the picker, the hidden file input, and a status line.
func (widget UploadMultiple) Edit(_ *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	const location = "widget.UploadMultiple.Edit"

	// RULE: The input's name is the host step's field name, which the widget cannot guess
	name := e.Options.GetString("name", value)

	if name == "" {
		return derp.Internal(location, "The upload-multiple widget requires a 'name' option", e.ID, e.Label)
	}

	remoteURL, err := widget.remoteURL(e, value)

	if err != nil {
		return derp.Wrap(err, location, "Unable to render url")
	}

	stagingURL, err := widget.stagingURL(e, value)

	if err != nil {
		return derp.Wrap(err, location, "Unable to render staging", e.ID, e.Label)
	}

	elementID := widget.elementID(e, value)
	accept := e.Options.GetString("accept", value)

	// The zone carries one of the two URLs, which is how the browser behavior tells a saved
	// object, whose files it changes in place, from a new one, whose files it stages
	b.Div().
		ID(elementID).
		Class("upload-multiple").
		Data("upload-multiple", "true").
		Data("name", name).
		Data("accept", accept).
		Data("url", remoteURL).
		Data("staging-url", stagingURL)

	if stagingURL == "" {
		widget.drawTileRow(b, elementID, remoteURL)
	} else {
		widget.drawStagingRow(b, elementID, name, stagingURL)
	}

	// RULE: Only the plus opens the picker, and it is a button because a button opens the
	// picker from the keyboard with no script. It follows the row so that it sits after
	// the last tile; the host's behavior hides it once the field is full.
	b.Button().ID(elementID+"-add").Type("button").Class("upload-multiple-add").Attr("title", "Add a file").Aria("label", "Add a file")
	b.I("bi", "bi-plus-lg").Close()
	b.Close()

	// RULE: Never `required`, and never `multiple`. The input is hidden, and a hidden invalid
	// control blocks a native submit with no visible cause; the host enforces the minimum.
	// The picker takes one file at a time, which is how the files are added.
	b.Input("file", name).
		ID(elementID+"-input").
		Attr("accept", accept).
		Attr("hidden", "hidden").
		Aria("label", e.Label).
		Aria("description", e.Description).
		Close()

	b.Div().ID(elementID + "-status").Class("upload-multiple-status").Role("status").Close()

	b.CloseAll()
	return nil
}

// remoteURL renders the url option, which names the host's route for the saved files.
// It is empty for an object that is not saved yet, because there is nothing to load.
func (widget UploadMultiple) remoteURL(e *form.Element, value any) (string, error) {

	if isNew(value) {
		return "", nil
	}

	// The url is an option template naming the object being edited,
	// so a broken one fails the render instead of drawing a dead row
	return e.Options.Evaluate("url", value)
}

// stagingURL renders the staging option, which names the host's route that holds a file
// until the object exists. It is empty for a saved object, whose own route takes the files.
func (widget UploadMultiple) stagingURL(e *form.Element, value any) (string, error) {

	const location = "widget.UploadMultiple.stagingURL"

	if !isNew(value) {
		return "", nil
	}

	result, err := e.Options.Evaluate("staging", value)

	if err != nil {
		return "", derp.Wrap(err, location, "Unable to render staging")
	}

	// RULE: A new object's own route does not exist yet, so without a staging route a
	// dropped file would have nowhere to go, and the zone would be a dead end
	if result == "" {
		return "", derp.Internal(location, "The upload-multiple widget requires a 'staging' option for an object that is not saved yet")
	}

	return result, nil
}

// elementID returns the element's own ID, or one made from its name, with dots turned to dashes.
func (widget UploadMultiple) elementID(e *form.Element, value any) string {

	result := e.ID

	if result == "" {
		result = e.Type + "-" + e.Options.GetString("name", value)
	}

	return strings.ReplaceAll(result, ".", "-")
}

// drawTileRow writes the empty tile row. With a URL, the row loads the host's own row
// in its place as soon as it is on the page.
func (widget UploadMultiple) drawTileRow(b *html.Builder, elementID string, remoteURL string) {

	if remoteURL == "" {
		b.Div().ID(elementID + "-tiles").Class("upload-multiple-tiles").Close()
		return
	}

	b.Div().ID(elementID+"-tiles").Class("upload-multiple-tiles").
		Attr("hx-get", remoteURL).
		Attr("hx-trigger", "load").
		Attr("hx-target", "this").
		Attr("hx-swap", "outerHTML").
		Attr("hx-push-url", "false").
		Close()
}

// drawStagingRow writes the empty tile row of a new object, wired to the staging route the
// way the host's row is wired to its own: each upload posts the file alone, and the tile the
// host answers with is appended, so the row is never replaced.
func (widget UploadMultiple) drawStagingRow(b *html.Builder, elementID string, name string, stagingURL string) {

	b.Div().ID(elementID+"-tiles").Class("upload-multiple-tiles").
		Attr("hx-post", stagingURL).
		Attr("hx-trigger", "upload-multiple-added").
		Attr("hx-params", name).
		Attr("hx-target", "this").
		Attr("hx-swap", "beforeend").
		Attr("hx-push-url", "false").
		Attr("hx-sync", "closest .upload-multiple:queue all").
		Close()
}

// isNew returns TRUE when the value is a record that has not been saved yet.
// Every benpate/data record answers IsNew; anything else counts as saved.
func isNew(value any) bool {

	if record, ok := value.(interface{ IsNew() bool }); ok {
		return record.IsNew()
	}

	return false
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For UploadMultiple widgets, labels are shown, so this always returns TRUE.
func (widget UploadMultiple) ShowLabels() bool {
	return true
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget UploadMultiple) ShowDescriptions() string {
	return "BOTTOM"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget.
// UploadMultiple widgets submit files, so this always returns "multipart/form-data".
func (widget UploadMultiple) Encoding(_ *form.Element) string {
	return "multipart/form-data"
}
