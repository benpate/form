package vocabulary

import (
	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/schema"
)

// LayoutVertical defines a standard top to bottom layout, including labels above every child item.
func LayoutVertical(library *form.Library) {

	library.Register("layout-vertical", func(form form.Form, schema *schema.Schema, value interface{}, b *html.Builder) error {

		var result error

		b.Div().Class("layout-vertical")

		if len(form.Label) > 0 {
			b.Div().Class("layout-vertical-label").InnerHTML(form.Label).Close()
		}

		b.Div().Class("layout-vertical-elements")

		for index, child := range form.Children {

			b.Div().Class("layout-vertical-element").Script(child.Show.HyperscriptRules())

			if child.Kind == "checkbox" {

				b.Label(child.ID)
				if err := child.Write(library, schema, value, b.SubTree()); err != nil {
					result = derp.Append(result, derp.Wrap(err, "form.widget.LayoutVertical", "Error rendering child", index, form))
				}
				b.Span().InnerHTML(child.Label).Close()
				b.Div().Class("text-sm", "gray40", "space-left").InnerHTML(child.Description).Close()
				b.Close()

			} else {
				b.Label(child.ID).InnerHTML(child.Label).Close()

				if err := child.Write(library, schema, value, b.SubTree()); err != nil {
					result = derp.Append(result, derp.Wrap(err, "form.widget.LayoutVertical", "Error rendering child", index, form))
				}

				if child.Description != "" {
					b.Div().Class("text-sm gray40").InnerHTML(child.Description).Close()
				}
			}

			b.Close()
		}

		b.CloseAll()

		return result
	})
}
