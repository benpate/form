# Form 🤔

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/form)
[![Version](https://img.shields.io/github/v/release/benpate/form?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/form/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/benpate/form/go.yml?style=flat-square)](https://github.com/benpate/form/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/form?style=flat-square)](https://goreportcard.com/report/github.com/benpate/form)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/form.svg?style=flat-square)](https://codecov.io/gh/benpate/form)

## JSON to HTML (Forms for Go)

This [Go module](https://golang.org) generates HTML forms using JSON configurations.  It is inspired by [JSON-Forms](https://jsonforms.io) but is not 100% compatible with that standard.

Forms can be defined directly in Go source code, or can be marshaled / unmarshalled from JSON.  See the section below for details on JSON marshaling.

## Form Definition
Following the design of JSON-Forms, form definitions are split into two parts: the data schema and the UI schema.

### Data Schema
Data schemas define the layout of the data that the form will display, relative to object in the application.  Schemas is created using the [Rosetta schema package](https://github.com/benpate/rosetta/tree/main/schema).  Please see that package for documentation on how to create a Data schema.

### UI Schema
UI Schemas define the UI elements that users interact with.  In this case, the HTML that will be delivered to the user's web browser.

### Example Code

```form := New(
````go
	schema.New(schema.Object{
		Properties: schema.ElementMap{
			"name":       schema.String{Required:true},
			"email":      schema.String{Format:"email"},
			"age":        schema.Integer{},
		},
	}),
	Element{
		Type: "test",
		Children: []Element{
			{Type: "text", Path: "name"},
			{Type: "text", Path: "age"},
			{Type: "text", Path: "email"},
		},
	},
)
```


## Form Elements
Each form is defined as a cascading tree of [form.Element](https://pkg.go.dev/github.com/benpate/form#Element) objects.  Each element defines standard properties along with an optional list of child elements

| Property    | Description                                                                                                          |
| ----------- | -------------------------------------------------------------------------------------------------------------------- |
| type        | String: Widget type to use when rendering this element.  Widgets are defined below                                   |
| label       | String: Label to put on the field. Usually renders above input elements in standard text.                            |
| description | String: Description for the field.  Usually renders below input elements  in small gray text                         |
| path        | String: Dot notation path to where the field value is stored                                                         |
| id          | String: DOM ID to use for this field.  Useful for some Javascript manipulations                                      |
| options     | Object: key/value pairs of non-standard options to pass to the widget element. Options are specific to each element. |
| children    | Array of Elements: an array of zero or more elements to render as "child" elements. See "Layouts" below.             |
| readOnly    | Boolean: If TRUE, then this element is rendered not as an input widget, but as a text value.                         |
|             |                                                                                                                      |


## Widgets
Form elements are rendered by widgets.  This project ships with a standard library of widgets in the [widgets package](https://pkg.go.dev/github.com/benpate/form/widget), but it is possible to register custom widgets by passing a [widget interface](https://pkg.go.dev/github.com/benpate/form#Widget) to the [`form.Use()`](https://pkg.go.dev/github.com/benpate/form#Use) function.

### type: "checkbox"
Documentation TBD

### type: "check-button"
This is a custom Hyperscript widget that displays a styled DIV that looks like a selectable button. Underneath is a simple checkbox that can be checked or not.  If the checkbox is checked, then the button appears selected.  Otherwise, the button appears in a default state.

| Option | Description                                                                                                  |
| ------ | ------------------------------------------------------------------------------------------------------------ |
| class  | [HTML5 class attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/class) |
| script | [Hyperscript script attribute](https://hyperscript.org/docs/#basics)                                         |

### type: "check-button-group"
This is a custom Hyperscript widget that uses an Option List to display a group of `check-button` widgets.  Option Lists are defined below, and can be pulled from an `option provider` or an `enum` defined in the element options object.

### type: "colorpicker"
This is a custom Hyperscript widget that displays an enhanced color picker input.  It uses the [HTML5 Color picker control](https://developer.mozilla.org/en-US/docs/Learn_web_development/Extensions/Forms/HTML5_input_types#color_picker_control "Color picker control") in addition to a scripted input element that displays the selected color and its hex code.

This widget does not take any optional arguments.

### type: "date"
This renders an [HTML5 date element](https://developer.mozilla.org/en-US/docs/Learn_web_development/Extensions/Forms/HTML5_input_types#date "date"), converting number and string values into a date.  This is the preferred method for adding/editing dates using this library

### type: datetime"
This renders an [HTML5 datetime-local element](https://developer.mozilla.org/en-US/docs/Learn_web_development/Extensions/Forms/HTML5_input_types#datetime-local "datetime-local"), converting number and string values into datetimes.  This is the preferred method for adding/editing date-times using this library.

### type: "heading"
This renders an HTML5 [H2 section heading](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements).  The element `Label` is used as the heading.  If an element `Description` is provided, then it is displayed beneath the `H2` in a standard `div` element.

### type: "hidden"
This renders an [HTML5 hidden input](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/hidden), with one option: 

If you want to set a fixed value for this input (instead of pulling its value from the Form's value object) you can use `options:{"value":"fixed value here"}` to set a fixed value for this hidden element.

### type: "html"
This renders custom HTML in the form, using the `Description` field as the HTML that is entered.

### type: "label"
This renders custom content in the form using one or both of the `Label` and `Description` fields.  

If a `Label` is provided, then it is entered into the HTML result as plain text. HTML values are escaped, if necessary.  

If a `Description` is provided, then it is entered into the HTML as HTML.

### type: "multiselect"
This renders is a custom Hyperscript widget that looks like a scrollable multiselect input.

### type: "password"
This renders a [HTML5 password](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/password) input, with the following additional options:

|              |                                                                                                                      |
| ------------ | -------------------------------------------------------------------------------------------------------------------- |
| autocomplete | [HTML5 autocomplete attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/autocomplete)  |
| autofocus    | [HTML5 autofocus attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/autofocus) |
| maxlength    | [HTML5 maxlength attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/maxlength)        |
| minlength    | [HTML5 minlength attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/minlength)        |
| pattern      | [HTML5 pattern attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/pattern)            |
| placeholder  | [HTML5 placeholder attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/placeholder)    |
| required     | [HTML5 required attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/required)          |
| style        | [HTML5 style attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/style)         |

### type: "radio"
Documentation TBD

### type: "select"
Renders an HTML5 select input.  The options available in the select box are dependent on the Option List (define below).  If the provided lookup code group is "writable" then an additional option `(Add New Value)` is added to the bottom of the list.

This widget allows the following additional options.

| Option   | Description                                                                                                                       |
| -------- | --------------------------------------------------------------------------------------------------------------------------------- |
| required | [HTML5 required attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/required)                       |
| focus    | [HTML5 autofocus attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/autofocus) (mislabeled) |

### type: "text"
If the form element type is `Text` then an HTML5 text input will be displayed.  There are a number of custom options for this element:

| Option      | Description                                                                                                                                      |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| autofocus   | [HTML5 autofocus attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/autofocus)                             |
| placeholder | [HTML5 placeholder attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/placeholder)                                |
| provider    | Option List provider (described below). Renders [HTML5  datalist](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/datalist) |
| script      | [Hyperscript script attribute](https://hyperscript.org/docs/#basics)                                                                             |
| spellcheck  | [HTML5 spellcheck attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/spellcheck)                           |
| style       | [HTML5 style attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/style)                                     |
| validator   | Remote validator URL for validating this widget from the server                                                                                  |

Additional options are available depending on the data type used in the schema element connected to this widget.

#### schema.Type = Integer 
If the associated schema is an `Integer` value, then the widget will render an [HTML5 Numeric field](https://developer.mozilla.org/en-US/docs/Learn_web_development/Extensions/Forms/HTML5_input_types#numeric_field) with a default step=1.  The following additional options are available:

| Option | Description                                                                                         |
| ------ | --------------------------------------------------------------------------------------------------- |
| min    | [HTML5 min attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/min)   |
| max    | [HTML5 max attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/max)   |
| step   | [HTML5 step attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/step) |

#### schema.Type = Number
If the associated schema is an `Number` value, then the widget will render an [HTML5 Numeric field](https://developer.mozilla.org/en-US/docs/Learn_web_development/Extensions/Forms/HTML5_input_types#numeric_field).  The following additional options are available:

| Option | Description                                                                                         |
| ------ | --------------------------------------------------------------------------------------------------- |
| min    | [HTML5 min attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/min)   |
| max    | [HTML5 max attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/max)   |
| step   | [HTML5 step attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/step) |

#### schema.Type = String
If the associated schema is an `String` value, then the widget will render a different HTML5 field depending on the schema.Type.

| Format         | Behavior                                                                                                                                     |
| -------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| date           | [Date field](https://developer.mozilla.org/en-US/docs/Learn_web_development/Extensions/Forms/HTML5_input_types#date)                         |
| datetime       | [DateTime-Local field](https://developer.mozilla.org/en-US/docs/Learn_web_development/Extensions/Forms/HTML5_input_types#datetime-local)     |
| email          | [Email address field](https://developer.mozilla.org/en-US/docs/Learn_web_development/Extensions/Forms/HTML5_input_types#email_address_field) |
| time           | [Time field](https://developer.mozilla.org/en-US/docs/Learn_web_development/Extensions/Forms/HTML5_input_types#time)                         |
| tel            | [Phone number field](https://developer.mozilla.org/en-US/docs/Learn_web_development/Extensions/Forms/HTML5_input_types#phone_number_field)   |
| text (default) | Plain text input                                                                                                                             |
| url            | [URL field](https://developer.mozilla.org/en-US/docs/Learn_web_development/Extensions/Forms/HTML5_input_types#url_field)                     |

In addition, the following options are available when rendering `String` values.

| Option    | Description                                                                                                   |
| --------- | ------------------------------------------------------------------------------------------------------------- |
| maxlength | [HTML5 maxlength attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/maxlength) |
| minlength | [HTML5 minlength attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/minlength) |
| pattern   | [HTML5 pattern attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/pattern)     |

### type: "textarea"
I the form element type is `Textarea` then an HTML5 textarea will be displayed.  There are a number of custom options for this element:

| Option    | Description                                                                                                              |
| --------- | ------------------------------------------------------------------------------------------------------------------------ |
| autofocus | [HTML5 autofocus attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/autofocus)     |
| rows      | [HTML5 rows attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/textarea#rows "rows")        |
| style     | [HTML5 style attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/style)             |
| minlength | [HTML5 minlength attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/minlength)            |
| maxlength | [HTML5 maxlength attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/maxlength)            |
| pattern   | [HTML5 pattern attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/pattern)                |
| required  | [HTML5 required attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/required)              |
| showLimit | If TRUE, then a Javascript counter will display the number of characters remaining before the maxlength limit is reached |

### type: "time"
This renders an[HTML5 time input](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/time).  It converts numbers and strings into a date-time value, then displays only the time portion in an input widget.

### type: "toggle"
Toggle is a custom HTML widget defined for this library.  It mimics the [iOS toggle control](https://developer.apple.com/design/human-interface-guidelines/toggles) in HTML.  It returns either a `true` or `false` value when the form is submitted.

| Option     | Value                                                          |
| ---------- | -------------------------------------------------------------- |
| text       | Text to appear next to the toggle button                       |
| true-text  | Text to appear next to the toggle button IF the value is TRUE  |
| false-text | Text to appear next to the toggle button IF the value is FALSE |


## Layouts
Layouts are a special kind of widget that contains other widgets.  Layouts are used as a general container for a set of form elements.  The child widgets inside of a layout are defined in the `children` property.

Layout widgets do not use any other custom options.

### Layout-Group
Displays child elements in vertical groups, one above the other.  Each child's `Label` is a heading above the remainder of its content.

### Layout- Horizontal
Displays child elements horizontally.  Each child's `Label` is placed above the widget.

This layout does not scale well in tight spaces and should be used sparingly.

### Layout-Tabs
Displays child elements as tabs inside a tab container.  Each child's `Label` is the label of the tab, and its `Description` is shown in text at the top of the tab panel.

Tabs are defined using [W3C-ARIA roles for tabs](https://www.w3.org/WAI/ARIA/apg/patterns/tabs/):  `tablist`, `tab`, and `tabpanel`

### Layout-Vertical
Displays child elements in a vertical column, placing each child's `Label` in plain text above the widget, and its `Description` below the widget in small gray text.

This is the most common layout for basic HTML5 forms.

## Option Lists
Certain widgets -- such as select boxes -- are able to display a list of options for the user to choose from.  Options are defined in a few different ways:

### Define Options Directly as an `enum`
Documentation TBD

### Importing Options from the Application use a `provider` 
Documentation TBD

## JSON Marshaling/Unmarshalling 
This package is intended to be used primarily by app developers to define forms in JSON configuration files.  So, extra care is taken to make JSON marshaling / unmarshalling work well. 

When building a form a JSON, all property names use [lower camel case](https://en.wikipedia.org/wiki/Camel_case), with the first letter of the property name as lower case.  If present, second and third words in the property name are capitalized

### JSON Example
Here is a sample form configuration in JSON.  This form is displayed in multiple tabs, adds/edits a number of values in the object (an "Album" in this case) and even has a place for file uploads.

```json
{
	"label": "Album",
	"type":"layout-tabs",
	"children": [
		{
			"type":"layout-vertical",
			"label": "General",
			"children": [
				{"type":"text", "path":"label", "label":"Album Name"},
				{"type":"select", "path":"data.license", "label":"License", "options":{"required":true, "provider":"bandwagon-album.licenses"}},
				{"type":"date", "path":"data.releaseDate", "label":"Release Date"},
				{"type":"upload", "path":"iconUrl", "label":"Album Art", "options":{"accept":"image/*", "filename":"data.imageFilename", "delete":"/{{.StreamID}}/delete-icon"}},
				{"type":"toggle", "path":"isFeatured", "options":{"true-text":"Featured (shows on home page)", "false-text":"Featured?"}}
			]
		},
		{
			"type":"layout-vertical",
			"label": "Metadata",
			"children": [
				{"type":"textarea", "path":"summary", "label":"Sidebar Notes", "description":"Notes appear on the side of the album page. Markdown is allowed.", "options":{"rows":8, "showLimit":true}},
				{"type":"textarea", "path":"data.tags", "label":"Tags", "description":"Enter #Hashtags separated by spaces."}
			]
		},
		{
			"type":"layout-vertical",
			"label": "Links",
			"children": [
				{"type":"text", "path":"data.links.AMAZON", "options":{"placeholder":"Amazon Music"}},
				{"type":"text", "path":"data.links.APPLE", "options":{"placeholder":"Apple Music"}},
				{"type":"text", "path":"data.links.BANDCAMP", "options":{"placeholder":"Bandcamp"}},
				{"type":"text", "path":"data.links.GOOGLE", "options":{"placeholder":"Google Play"}},
				{"type":"text", "path":"data.links.SOUNDCLOUD", "options":{"placeholder":"Soundcloud"}},
				{"type":"text", "path":"data.links.SPOTIFY", "options":{"placeholder":"Spotify"}},
				{"type":"text", "path":"data.links.TIDAL", "options":{"placeholder":"Tidal"}},
				{"type":"text", "path":"data.links.YOUTUBE", "options":{"placeholder":"YouTube Music"}},
				{"type":"text", "path":"data.links.OTHER1", "options":{"placeholder":"Other"}},
				{"type":"text", "path":"data.links.OTHER2", "options":{"placeholder":"Other"}},
				{"type":"text", "path":"data.links.OTHER3", "options":{"placeholder":"Other"}}
			]
		},
		{
			"type":"layout-vertical",
			"label": "Colors",
			"children": [
				{"type":"colorpicker", "path":"data.color.body", "label":"Window Background"},
				{"type":"colorpicker", "path":"data.color.page", "label":"Page Background"},
				{"type":"colorpicker", "path":"data.color.button", "label":"Links and Buttons"}
			]
		}
	]
}
```


# Project Status

This project is a work-in-progress.  It is being used in production on large websites, but it is still under active development and is subject to change without notice.

If you're looking for a form library in Go, you should probably use something else that has a better stability guarantee.

With that said, if you have an idea for making this library better, send in a pull request.  We're all in this together! 🤔
