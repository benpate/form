package widget

import (
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
)

func drawLayout(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder, alignment string, edit bool) error {

	const location = "form.drawLayout"
	var result error

	b.Div().Class("layout", "layout-"+alignment)

	if len(e.Label) > 0 {
		b.Div().Class("layout-title").InnerHTML(e.Label).Close()
	}

	if len(e.Description) > 0 {
		b.Div().Class("layout-description").InnerHTML(e.Description).Close()
	}

	b.Div().Class("layout-elements")

	for index := range e.Children {

		child := e.Children[index]

		// If there is a "show-if-option" then look in the form.Options for a true/false
		if showIfToken := child.Options.GetString("show-if-option"); showIfToken != "" {
			if showValue := f.OptionBool(showIfToken); !showValue {
				continue
			}
		}

		// If there is a "hide-if-option" then look in the form.Options for a true/false
		if hideIfToken := child.Options.GetString("hide-if-option"); hideIfToken != "" {
			if hideValue := f.OptionBool(hideIfToken); hideValue {
				continue
			}
		}

		// Set default ID if not present
		if child.ID == "" {
			child.ID = strings.ReplaceAll(child.Path, ".", "_") + "_" + child.Type
		}

		// Locate the widget that will draw this child
		widget, err := child.Widget()
		if err != nil {
			return derp.Wrap(err, location, "Unable to find selected widget", index, child)
		}

		var container *html.Element

		// All elements (except hidden and toggle-group) get wrapped in a div
		if (child.Type != "hidden") && (child.Type != "toggle-group") {
			container = b.Div()

			container.Class("layout-element", "layout-"+alignment+"-element")

			if showIf := child.Options.GetString("show-if"); showIf != "" {
				container.Data("script", "install showIf(condition:'"+showIf+"')")
			}

			container.EndBracket()

			if widget.ShowLabels() {
				b.Label(child.ID).Aria("hidden", "true").InnerText(child.Label).Close()
			}

			if child.Description != "" {
				if widget.ShowDescriptions() == "TOP" {
					b.Div().Aria("hidden", "true").Class("text-sm gray40").InnerHTML(child.Description).Close()
				}
			}
		}

		// Draw the edit or view version of this element
		if edit {
			if err := child.Edit(f, provider, value, b.SubTree()); err != nil {
				return derp.Wrap(err, location, "Unable to draw child (edit)", e, index, child)
			}
		} else {
			if err := child.View(f, provider, value, b.SubTree()); err != nil {
				return derp.Wrap(err, location, "Unable to draw child (view)", e, index, child)
			}
		}

		// If there's a description on this element, draw it here
		if child.Description != "" {
			if widget.ShowDescriptions() == "BOTTOM" {
				b.Div().Aria("hidden", "true").Class("text-sm gray40").InnerHTML(child.Description).Close()
			}
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
func collectEncoding(children []form.Element) string {

	for _, child := range children {

		if widget, err := child.Widget(); err == nil {
			if encoding := widget.Encoding(&child); encoding != "" {
				return encoding
			}
		}
	}

	return ""
}
