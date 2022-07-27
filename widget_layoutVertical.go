package form

import (
	"github.com/benpate/derp"
	"github.com/segmentio/ksuid"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func init() {
	Register("layout-vertical", HTMLLayoutVertical)
}

func HTMLLayoutVertical(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder) error {

	var result error

	b.Div().Class("layout-vertical")

	if len(element.Label) > 0 {
		b.Div().Class("layout-title").InnerHTML(element.Label).Close()
	}

	b.Div().Class("layout-vertical-elements")

	for index := range element.Children {

		// Default ID for this child element
		if element.Children[index].ID == "" {
			element.Children[index].ID = ksuid.New().String()
		}

		child := element.Children[index]

		// Special case for hidden fields (don't include element wrappers)
		if child.Type == "hidden" {
			if err := child.WriteHTML(schema, lookupProvider, value, b.SubTree()); err != nil {
				result = derp.Append(result, derp.Wrap(err, "form.widget.LayoutVertical", "Error rendering hidden field", element, index, child))
			}
			continue
		}

		b.Div().Class("layout-vertical-element")

		if child.Type == "checkbox" {

			b.Label(child.ID)
			if err := child.WriteHTML(schema, lookupProvider, value, b.SubTree()); err != nil {
				result = derp.Append(result, derp.Wrap(err, "form.widget.LayoutVertical", "Error rendering child", element, index, child))
			}
			b.Span().InnerHTML(child.Label).Close()
			b.Div().Class("text-sm", "gray40", "space-left").InnerHTML(child.Description).Close()
			b.Close()

		} else {
			b.Label(child.ID).InnerHTML(child.Label).Close()

			if err := child.WriteHTML(schema, lookupProvider, value, b.SubTree()); err != nil {
				result = derp.Append(result, derp.Wrap(err, "form.widget.LayoutVertical", "Error rendering child", element, index, child))
			}

			if child.Description != "" {
				b.Div().Class("text-sm gray40").InnerHTML(child.Description).Close()
			}
		}

		b.Close()
	}

	b.CloseAll()

	return result
}
