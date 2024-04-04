package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("upload", WidgetUpload{})
}

type WidgetUpload struct{}

func (WidgetUpload) View(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	valueString := element.GetString(value, s)

	if valueString == "" {
		valueString = "N/A"
	}

	b.Div().Class("layout-value", element.Options.GetString("class")).InnerText(valueString).Close()
	return nil
}

func (WidgetUpload) Edit(element *Element, s *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return WidgetUpload{}.View(element, s, lookupProvider, value, b)
	}

	elementID := element.ID

	if elementID == "" {
		elementID = "upload-" + element.Path
	}

	if valueString := element.GetString(value, s); valueString != "" {

		filename := valueString
		if filenamePath := element.Options.GetString("filename"); filenamePath != "" {
			if value, err := s.Get(value, filenamePath); err == nil {
				filename = convert.String(value)
			}
		}

		if filename != "" {
			b.Span().InnerText(filename).Close()
		}
	}

	multiple := iif(element.Options.GetBool("multiple"), "multiple", "")

	b.Input("file", element.Path).ID(elementID).
		Attr("accept", element.Options.GetString("accept")).
		Attr("multiple", multiple).
		Close()

	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (WidgetUpload) ShowLabels() bool {
	return true
}
