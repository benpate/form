package form

import (
	"github.com/benpate/derp"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("layout-horizontal", HTMLLayoutHorizontal)
}

func HTMLLayoutHorizontal(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	var result error

	b.Div().Class("layout-horizontal")
	if element.Label != "" {
		b.Div().Class("layout-horizontal-label").InnerHTML(element.Label).Close()
	}
	b.Div().Class("layout-horizontal-elements")

	for index, child := range element.Children {

		// Special case for hidden fields (don't include element wrappers)
		if child.Type == "hidden" {
			if err := child.WriteHTML(schema, lookupProvider, value, b.SubTree()); err != nil {
				result = derp.Append(result, derp.Wrap(err, "form.widget.LayoutVertical", "Error rendering hidden field", element, index, child))
			}
			continue
		}

		b.Div().Class("layout-horizontal-element")

		if element.Options.GetBool("show-labels") {
			b.Div().Class("label").InnerHTML(child.Label).Close()
		}

		if err := child.WriteHTML(schema, lookupProvider, value, b.SubTree()); err != nil {
			result = derp.Wrap(err, "form.widget.LayoutHorizontal", "Error rendering child", element, index, child)
		}

		b.Close()
	}

	b.CloseAll()

	return result
}
