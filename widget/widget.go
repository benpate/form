package widget

import (
	"github.com/benpate/form"
)

func UseAll() {
	form.Use("checkbox", Checkbox{})
	form.Use("colorpicker", Colorpicker{})
	form.Use("datepicker", DatePicker{})
	form.Use("heading", Heading{})
	form.Use("hidden", Hidden{})
	form.Use("label", Label{})
	form.Use("layout-group", LayoutGroup{})
	form.Use("layout-horizontal", LayoutHorizontal{})
	form.Use("layout-tabs", LayoutTabs{})
	form.Use("layout-vertical", LayoutVertical{})
	form.Use("multiselect", Multiselect{})
	form.Use("password", Password{})
	form.Use("radio", Radio{})
	form.Use("select", Select{})
	form.Use("text", Text{})
	form.Use("textarea", TextArea{})
	form.Use("toggle", Toggle{})
	form.Use("upload", Upload{})
	form.Use("wysiwyg", WYSIWYG{})
}
