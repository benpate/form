package form

import (
	"strings"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("upload", WidgetUpload{})
}

type WidgetUpload struct{}

func (w WidgetUpload) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := element.GetString(value, s)

	if valueString == "" {
		valueString = "N/A"
	}

	b.Div().Class("layout-value", element.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

func (w WidgetUpload) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetUpload{}.View(element, s, lookupProvider, value, b)
	}

	elementID := element.ID

	if elementID == "" {
		elementID = "upload-" + element.Path
	}

	w.preview(element, s, value, b.SubTree())

	multiple := iif(element.Options.GetBool("multiple"), "multiple", "")
	b.Input("file", element.Path).ID(elementID).
		Attr("accept", element.Options.GetString("accept")).
		Attr("multiple", multiple).
		Close()

	return nil
}

func (w WidgetUpload) preview(element *Element, s *schema.Schema, value any, b *html.Builder) {

	// Get the URL for the uploaded file
	valueString := element.GetString(value, s)

	if valueString == "" {
		return
	}

	// Different file types are displayed differently
	accept := element.Options.GetString("accept")
	acceptType, _, _ := strings.Cut(accept, "/")

	switch acceptType {

	// Image preview (128px square)
	case "image":
		b.Img(valueString).Style("border:solid 1px black", "max-height:128px", "max-width:128px").Close()

	// Audio previoew (with controls)
	case "audio":
		b.Audio().Attr("controls", "true")
		b.Source().Src(valueString).Close()
		b.Close()

	// All other files are displayed as a link
	default:
		b.A(valueString).InnerText(valueString).Close()
	}
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (w WidgetUpload) ShowLabels() bool {
	return true
}
