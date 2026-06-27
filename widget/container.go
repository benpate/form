package widget

import (
	"github.com/benpate/derp"
	"github.com/benpate/form"
	"github.com/benpate/html"
)

// Container is a widget that wraps its child elements in a styled <div>,
// optionally hiding children based on form options.
type Container struct{}

// View generates the read-only HTML for this container and its children.
func (Container) View(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {
	const location = "form.Container.View"
	var result error

	classes := e.Options.GetSliceOfString("class")
	styles := e.Options.GetSliceOfString("style")

	b.Div().Class(classes...).Style(styles...)

	for index, child := range e.Children {

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

		// Get the widget for this child
		widget, err := child.Widget()

		if err != nil {
			return derp.Wrap(err, location, "Unable to draw child", index, child)
		}

		// Draw the 'view' version of this element
		if err := widget.View(f, &child, provider, value, b.SubTree()); err != nil {
			return derp.Wrap(err, location, "Unable to draw child", e, index, child)
		}
	}

	b.CloseAll()

	return result
}

// Edit generates the editable HTML for this container and its children.
func (Container) Edit(f *form.Form, e *form.Element, provider form.LookupProvider, value any, b *html.Builder) error {

	const location = "form.Container.Edit"
	var result error

	classes := e.Options.GetSliceOfString("class")
	styles := e.Options.GetSliceOfString("style")

	b.Div().Class(classes...).Style(styles...)

	for index, child := range e.Children {

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

		// Get the widget for this child
		widget, err := child.Widget()

		if err != nil {
			return derp.Wrap(err, location, "Unable to draw child", index, child)
		}

		// Draw the 'edit' version of this element
		if err := widget.Edit(f, &child, provider, value, b.SubTree()); err != nil {
			return derp.Wrap(err, location, "Unable to draw child", e, index, child)
		}
	}

	b.CloseAll()

	return result
}

/***********************************
 * Widget Metadata
 ***********************************/

// ShowLabels is a part of the Widget interface.
// It returns TRUE if this widget requires labels to be displayed around it.
// For Container widgets, labels are not shown, so this always returns FALSE.
func (Container) ShowLabels() bool {
	return false
}

// ShowDescriptions is a part of the Widget interface.
// It returns the position of the description for this widget,
// which is either "TOP", "BOTTOM", or "NONE".
func (widget Container) ShowDescriptions() string {
	return "TOP"
}

// Encoding is a part of the Widget interface.
// It returns the encoding type for this widget,
// which is the combined encoding of all child elements.
func (widget Container) Encoding(element *form.Element) string {
	return collectEncoding(element.Children)
}
