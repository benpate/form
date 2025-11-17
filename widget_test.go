package form

import (
	"github.com/benpate/html"
)

type testWidget struct{}

func (w testWidget) View(_ *Form, element *Element, _ LookupProvider, _ any, b *html.Builder) error {
	b.Empty("widget-view").Attr("name", element.Path)
	return nil
}

func (w testWidget) Edit(_ *Form, element *Element, _ LookupProvider, _ any, b *html.Builder) error {
	b.Empty("widget-edit").Attr("name", element.Path)
	return nil
}

func (w testWidget) ShowLabels() bool {
	return false
}

func (w testWidget) Encoding(_ *Element) string {
	return ""
}

func useTestWidget() {
	Use("test", testWidget{})
}
