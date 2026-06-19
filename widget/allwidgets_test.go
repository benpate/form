package widget

import (
	"net/url"
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

// widgetData is a sample object that exercises the various schema paths.
func widgetData() mapof.Any {
	return mapof.Any{
		"username": "johnconnor",
		"name":     "John Connor",
		"email":    "john@connor.mil",
		"age":      42,
		"human":    true,
		"distance": 12.5,
		"color":    "Red",
		"tags":     []string{"pretty", "please"},
		"terms":    true,
		"other":    "some text",
		"ology": mapof.Any{
			"biology": "I am a biological human",
		},
	}
}

// drawWidget renders a single element through both the Editor and Viewer and
// confirms neither path errors.
func drawWidget(t *testing.T, element form.Element) (string, string) {
	t.Helper()

	UseAll()
	f := form.New(getTestSchema(), element)
	provider := testLookupProvider{}
	data := widgetData()

	editor, err := f.Editor(data, provider)
	require.NoError(t, err, "Editor for %q", element.Type)

	viewer, err := f.Viewer(data, provider)
	require.NoError(t, err, "Viewer for %q", element.Type)

	return editor, viewer
}

func TestWidget_InputWidgets(t *testing.T) {

	inputs := []form.Element{
		{Type: "text", Path: "name"},
		{Type: "textarea", Path: "ology.biology"},
		{Type: "password", Path: "other", Label: "Pass"},
		{Type: "hidden", Path: "name"},
		{Type: "toggle", Path: "human", Label: "Human"},
		{Type: "checkbox", Path: "human", Label: "Accept"},
		{Type: "colorpicker", Path: "color"},
		{Type: "date", Path: "other"},
		{Type: "datetime", Path: "other"},
		{Type: "time", Path: "other"},
		{Type: "upload", Path: "other"},
		{Type: "wysiwyg", Path: "ology.biology"},
	}

	for _, element := range inputs {
		editor, _ := drawWidget(t, element)
		require.NotEmpty(t, editor, "Editor output for %q should not be empty", element.Type)
	}
}

func TestWidget_LookupWidgets(t *testing.T) {

	// Widgets that render a set of options from a lookup provider / schema enum.
	lookups := []form.Element{
		{Type: "select", Path: "color"},
		{Type: "select-group", Path: "color", Options: mapof.Any{"provider": "test", "children": "other"}},
		{Type: "radio", Path: "color"},
		{Type: "radio-button-group", Path: "color"},
		{Type: "radio-button-group-horizontal", Path: "color"},
		{Type: "radio-colors", Path: "color"},
		{Type: "multiselect", Path: "tags", Options: mapof.Any{"provider": "test"}},
		{Type: "check-button", Path: "color", Options: mapof.Any{"provider": "test"}},
		{Type: "check-button-group", Path: "tags", Options: mapof.Any{"provider": "test"}},
	}

	for _, element := range lookups {
		drawWidget(t, element)
	}
}

func TestWidget_DisplayWidgets(t *testing.T) {

	displays := []form.Element{
		{Type: "heading", Label: "My Heading", Description: "A description"},
		{Type: "label", Label: "My Label", Description: "Help text"},
		{Type: "html", Description: "<b>raw html</b>"},
		{Type: "html-remote", Options: mapof.Any{"url": "/remote/{{.name}}"}},
	}

	for _, element := range displays {
		drawWidget(t, element)
	}
}

func TestWidget_LayoutWidgets(t *testing.T) {

	children := []form.Element{
		{Type: "text", Path: "name", Label: "Name"},
		{Type: "text", Path: "email", Label: "Email"},
	}

	layouts := []form.Element{
		{Type: "container", Children: children},
		{Type: "layout-group", Children: children},
		{Type: "layout-vertical", Children: children},
		{Type: "layout-horizontal", Children: children},
	}

	for _, element := range layouts {
		editor, _ := drawWidget(t, element)
		require.NotEmpty(t, editor, "Layout %q should produce output", element.Type)
	}

	// layout-tabs renders only in Edit mode (View is an intentional no-op),
	// so it is driven separately without the non-empty Viewer assertion.
	drawWidget(t, form.Element{Type: "layout-tabs", Children: children})
}

func TestWidget_HTMLRemote_InvalidTemplate(t *testing.T) {

	UseAll()
	f := form.New(getTestSchema(), form.Element{
		Type:    "html-remote",
		Options: mapof.Any{"url": "{{.unterminated"},
	})

	// A malformed URL template produces an error
	_, err := f.Editor(widgetData(), nil)
	require.Error(t, err)
}

func TestWidget_Place(t *testing.T) {

	UseAll()

	s := placeSchema()
	provider := testLookupProvider{}

	element := form.Element{Type: "place", Path: "location", ID: "loc"}
	f := form.New(s, element)

	value := mapof.Any{
		"location": mapof.Any{
			"name":      "Eiffel Tower",
			"formatted": "Paris, France",
			"latitude":  "48.8584",
			"longitude": "2.2945",
		},
	}

	editor, err := f.Editor(value, provider)
	require.NoError(t, err)
	require.NotEmpty(t, editor)

	viewer, err := f.Viewer(value, provider)
	require.NoError(t, err)
	require.NotEmpty(t, viewer)
}

func TestWidget_Place_SetURLValue(t *testing.T) {

	UseAll()

	s := placeSchema()
	f := form.New(s, form.Element{Type: "place", Path: "location", ID: "loc"})

	object := mapof.Any{}
	values := url.Values{
		"location.formatted": []string{"Paris, France"},
		"location.latitude":  []string{"48.8584"},
		"location.longitude": []string{"2.2945"},
	}

	require.NoError(t, f.SetURLValues(&object, values, nil))

	location := object.GetMapOfAny("location")
	require.Equal(t, "Paris, France", location.GetString("formatted"))
	require.Equal(t, "48.8584", location.GetString("latitude"))
	require.Equal(t, "2.2945", location.GetString("longitude"))
}

// placeSchema returns a schema with a "location" object suitable for the Place widget.
func placeSchema() schema.Schema {
	return schema.New(schema.Object{
		Properties: schema.ElementMap{
			"location": schema.Object{
				Properties: schema.ElementMap{
					"name":      schema.String{},
					"formatted": schema.String{},
					"latitude":  schema.String{},
					"longitude": schema.String{},
				},
			},
		},
	})
}

func TestWidget_Metadata(t *testing.T) {

	// Every registered widget must implement the metadata methods without panicking.
	UseAll()

	types := []string{
		"checkbox", "check-button", "check-button-group", "colorpicker", "container",
		"date", "datetime", "heading", "html", "html-remote", "hidden", "label",
		"layout-group", "layout-horizontal", "layout-tabs", "layout-vertical",
		"multiselect", "password", "place", "radio", "radio-button-group",
		"radio-button-group-horizontal", "radio-colors", "select", "select-group",
		"text", "textarea", "time", "toggle", "upload", "wysiwyg",
	}

	for _, widgetType := range types {
		element := form.Element{Type: widgetType}
		widget, err := element.Widget()
		require.NoError(t, err, widgetType)

		// Each metadata method must return without panicking
		_ = widget.ShowLabels()
		require.NotEmpty(t, widget.ShowDescriptions(), widgetType) // "TOP"/"BOTTOM"/"NONE"
		_ = widget.Encoding(&element)
	}
}
