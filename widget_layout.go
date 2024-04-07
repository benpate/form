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
		b.Div().Class("layout-title").InnerText(element.Label).Close()
	}

	if len(element.Description) > 0 {
		b.Div().Class("layout-description").InnerText(element.Description).Close()
	}

	b.Div().Class("layout-" + alignment + "-elements")

	for index := range element.Children {

		child := element.Children[index]

		widget, err := child.Widget()
		if err != nil {
			return derp.Wrap(err, location, "Error rendering child", index, child)
		}

		var container *html.Element

		// All elements (except hidden) get wrapped in a div
		if child.Type != "hidden" {
			container = b.Div()

			container.Class("layout-element", "layout-"+alignment+"-element")

			if showIf := child.Options.GetString("show-if"); showIf != "" {
				container.Data("script", "install showIf(condition:'"+showIf+"')")
			}

			container.EndBracket()

			if widget.ShowLabels() {
				b.Label(child.ID).InnerText(child.Label).Close()
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
			b.Div().Class("text-sm gray40").InnerText(child.Description).Close()
		}

		// Close the DIV wrapper from above (if applicable)
		if container != nil {
			container.Close()
		}
	}

	b.CloseAll()

	return result
}

// collectEncoding returns the first non-empty encoding found in a slice of child elements
func collectEncoding(children []Element) string {

	for _, child := range children {

		if widget, err := child.Widget(); err == nil {
			if encoding := widget.Encoding(&child); encoding != "" {
				return encoding
			}
		}
	}

	return ""
}
