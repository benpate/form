// Package widget provides the built-in form widgets (text inputs, selects,
// layouts, and more) that render and edit individual form Elements. Call
// UseAll to register every widget with the form package before drawing a form.
package widget

import (
	"github.com/benpate/form"
)

// UseAll registers every built-in widget into the form package's widget registry.
func UseAll() {
	form.Use("checkbox", Checkbox{})
	form.Use("check-button", CheckButton{})
	form.Use("check-button-group", CheckButtonGroup{})
	form.Use("colorpicker", Colorpicker{})
	form.Use("container", Container{})
	form.Use("date", DatePicker{})
	form.Use("datetime", DateTimePicker{})
	form.Use("heading", Heading{})
	form.Use("html", HTML{})
	form.Use("html-remote", HTMLRemote{})
	form.Use("hidden", Hidden{})
	form.Use("label", Label{})
	form.Use("layout-group", LayoutGroup{})
	form.Use("layout-horizontal", LayoutHorizontal{})
	form.Use("layout-tabs", LayoutTabs{})
	form.Use("layout-vertical", LayoutVertical{})
	form.Use("multiselect", Multiselect{})
	form.Use("password", Password{})
	form.Use("place", Place{})
	form.Use("radio", Radio{})
	form.Use("radio-button-group", RadioButtonGroup{})
	form.Use("radio-button-group-horizontal", RadioButtonGroupHorizontal{})
	form.Use("radio-colors", RadioColors{})
	form.Use("select", Select{})
	form.Use("select-group", SelectGroup{})
	form.Use("text", Text{})
	form.Use("textarea", TextArea{})
	form.Use("time", TimePicker{})
	form.Use("toggle", Toggle{})
	form.Use("upload", Upload{})
	form.Use("wysiwyg", WYSIWYG{})
}
