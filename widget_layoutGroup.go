package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
	"github.com/segmentio/ksuid"

	"github.com/benpate/derp"
)

func init() {
	Register("layoutGroup", HTMLLayoutGroup)
}

func HTMLLayoutGroup(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	var result error

	b.Div().Class("layout-group")

	if element.Label != "" {
		b.Div().Class("layout-group-label").InnerHTML(element.Label).Close()
	}

	b.Div().Class("layout-group-elements")

	for index := range element.Children {

		// Default ID for this child element
		if element.Children[index].ID == "" {
			element.Children[index].ID = ksuid.New().String()
		}

		child := element.Children[index]

		tag := b.Div().Class("layout-group-element")

		if err := child.WriteHTML(schema, lookupProvider, value, b.SubTree()); err != nil {
			result = derp.Wrap(err, "form.widget.LayoutGroup", "Error rendering child", index, child)
		}

		tag.Close()
	}

	b.CloseAll()

	return result
}
