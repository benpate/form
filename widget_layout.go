package form

import (
	"github.com/benpate/derp"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

func drawLayout(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, b *html.Builder, alignment string, edit bool) error {

	const location = "form.drawLayout"
	var result error

	b.Div().Class("layout", "layout-"+alignment)

	if len(element.Label) > 0 {
		b.Div().Class("layout-title").InnerHTML(element.Label).Close()
	}

	if len(element.Description) > 0 {
		b.Div().Class("layout-description").InnerHTML(element.Description).Close()
	}

	b.Div().Class("layout-" + alignment + "-elements")

	for index := range element.Children {

		child := element.Children[index]

		widget, err := child.Widget()
		if err != nil {
			return derp.Wrap(err, location, "Error rendering child", index, child)
		}

		// All elements (except hidden) get wrapped in a div
		if child.Type != "hidden" {
			b.Div().Class("layout-" + alignment + "-element")

			if widget.ShowLabels() {
				b.Label(child.ID).InnerHTML(child.Label).Close()
			}
		}

		// Draw the edit or view version of this element
		if edit {
			if err := widget.Edit(&child, schema, lookupProvider, value, b.SubTree()); err != nil {
				return derp.Wrap(err, location, "Error rendering child (edit)", element, index, child)
			}
		} else {
			if err := widget.View(&child, schema, lookupProvider, value, b.SubTree()); err != nil {
				return derp.Wrap(err, location, "Error rendering child (view)", element, index, child)
			}
		}

		// If there's a description on this element, draw it here
		if widget.ShowLabels() && (child.Description != "") {
			b.Div().Class("text-sm gray40").InnerHTML(child.Description).Close()
		}

		// Close the DIV wrapper from above
		if child.Type != "hidden" {
			b.Close()
		}
	}

	b.CloseAll()

	return result
}
