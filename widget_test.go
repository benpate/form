package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

type testWidget struct{}

func (w testWidget) View(element *Element, s *schema.Schema, provider LookupProvider, value any, b *html.Builder) error {
	b.Empty("widget-view").Attr("name", element.Path)
	return nil
}

func (w testWidget) Edit(element *Element, s *schema.Schema, provider LookupProvider, value any, b *html.Builder) error {
	b.Empty("widget-edit").Attr("name", element.Path)
	return nil
}

func (w testWidget) ShowLabels() bool {
	return false
}

func (w testWidget) Encoding(element *Element) string {
	return ""
}

func useTestWidget() {
	Use("test", testWidget{})
}
