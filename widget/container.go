package widget

import (
	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
)

type Container struct{}

func (Container) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	const location = "form.Container.View"
	var result error

	classes := e.Options.GetSliceOfString("class")
	styles := e.Options.GetSliceOfString("style")

	b.Div().Class(classes...).Style(styles...)

	for index, child := range e.Children {

		// Get the widget for this child
		widget, err := child.Widget()

		if err != nil {
			return derp.Wrap(err, location, "Error rendering child", index, child)
		}

		// Draw the 'view' version of this element
		if err := widget.View(f, &child, provider, value, b.SubTree()); err != nil {
			return derp.Wrap(err, location, "Error rendering child", e, index, child)
		}
	}

	b.CloseAll()

	return result
}

func (Container) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	const location = "form.Container.Edit"
	var result error

	classes := e.Options.GetSliceOfString("class")
	styles := e.Options.GetSliceOfString("style")

	b.Div().Class(classes...).Style(styles...)

	for index, child := range e.Children {

		// Get the widget for this child
		widget, err := child.Widget()

		if err != nil {
			return derp.Wrap(err, location, "Error rendering child", index, child)
		}

		// Draw the 'edit' version of this element
		if err := widget.Edit(f, &child, provider, value, b.SubTree()); err != nil {
			return derp.Wrap(err, location, "Error rendering child", e, index, child)
		}
	}

	b.CloseAll()

	return result
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (Container) ShowLabels() bool {
	return false
}

func (widget Container) Encoding(element *form.Element) string {
	return collectEncoding(element.Children)
}
