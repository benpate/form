package widget

import (
	"strings"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

type Upload struct{}

func (widget Upload) View(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := element.GetString(value, s)

	if valueString == "" {
		valueString = "N/A"
	}

	b.Div().Class("layout-value", element.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

func (widget Upload) Edit(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return Upload{}.View(element, s, nil, value, b)
	}

	elementID := element.ID

	if elementID == "" {
		elementID = element.Path + ".upload"
	}

	widget.preview(element, s, value, b.SubTree())

	multiple := iif(element.Options.GetBool("multiple"), "multiple", "")
	b.Input("file", element.Path).ID(elementID).
		Attr("accept", element.Options.GetString("accept")).
		Attr("multiple", multiple).
		Aria("label", element.Label).
		Aria("description", element.Description).
		TabIndex("0").
		Close()

	return nil
}

func (widget Upload) preview(element *form.Element, s *schema.Schema, value any, b *html.Builder) {

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
		b.Div().Class("pos-relative", "width-128").Style("border:solid 1px black")
		b.Img(valueString).Style("max-width:128px", "max-height:128px").Close()

		if deleteLink := element.Options.GetString("delete"); deleteLink != "" {
			b.Button().
				Style("position:absolute", "top:4px", "right:4px").
				Class("text-xs").
				Attr("hx-post", deleteLink).
				Attr("hx-confirm", "Delete this file?").
				Attr("script", "on htmx:afterRequest log me then log my parentNode then remove my parentNode").
				Aria("label", "Delete").
				InnerText("X").
				Close()
		}
		b.Close()

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

func (widget Upload) ShowLabels() bool {
	return true
}

func (widget Upload) Encoding(_ *form.Element) string {
	return "multipart/form-data"
}
