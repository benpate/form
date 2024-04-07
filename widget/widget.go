package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/form/widget"
)

func Use() {
	form.Use("checkbox", widget.Checkbox{})
	form.Use("heading", widget.Heading{})
	form.Use("hidden", widget.Hidden{})
	form.Use("label", widget.Label{})
	form.Use("layout-group", widget.LayoutGroup{})
	form.Use("layout-horizontal", widget.LayoutHorizontal{})
	form.Use("layout-tabs", widget.LayoutTabs{})
	form.Use("layout-vertical", widget.LayoutVertical{})
	form.Use("multiselect", widget.Multiselect{})
	form.Use("radio", widget.Radio{})
	form.Use("select", widget.Select{})
	form.Use("text", widget.Text{})
	form.Use("textarea", widget.Textarea{})
	form.Use("toggle", widget.Toggle{})
	form.Use("upload", widget.Upload{})
	form.Use("wysiwyg", widget.WYSIWYG{})
}
