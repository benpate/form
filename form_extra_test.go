package form

import (
	"testing"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

// testForm builds a working Form (with the registered "test" widget) over a
// simple single-string schema.
func testForm(t *testing.T, options ...string) (Form, any) {
	t.Helper()
	useTestWidget()

	s := schema.New(schema.Object{
		Properties: schema.ElementMap{
			"name": schema.String{},
		},
	})

	element := Element{Type: "test", Path: "name"}
	value := mapof.Any{"name": "Sarah"}

	return New(s, element, options...), value
}

func TestForm_Editor(t *testing.T) {
	form, value := testForm(t)

	result, err := form.Editor(value, nil)
	require.NoError(t, err)
	require.Equal(t, `<widget-edit name="name">`, result)
}

func TestForm_Viewer(t *testing.T) {
	form, value := testForm(t)

	result, err := form.Viewer(value, nil)
	require.NoError(t, err)
	require.Equal(t, `<widget-view name="name">`, result)
}

func TestForm_StandaloneEditorViewer(t *testing.T) {
	useTestWidget()
	s := schema.New(schema.Object{Properties: schema.ElementMap{"name": schema.String{}}})
	element := Element{Type: "test", Path: "name"}
	value := mapof.Any{"name": "x"}

	edited, err := Editor(s, element, value, nil)
	require.NoError(t, err)
	require.Equal(t, `<widget-edit name="name">`, edited)

	viewed, err := Viewer(s, element, value, nil)
	require.NoError(t, err)
	require.Equal(t, `<widget-view name="name">`, viewed)
}

func TestForm_Encoding(t *testing.T) {
	form, _ := testForm(t)
	// testWidget.Encoding returns ""
	require.Equal(t, "", form.Encoding())
}

func TestForm_Options(t *testing.T) {

	form, _ := testForm(t, "title:My Form", "columns: 3", "compact: true")

	require.Equal(t, "My Form", form.OptionString("title"))
	require.Equal(t, "", form.OptionString("missing"))

	require.Equal(t, 3, form.OptionInt("columns"))
	require.Equal(t, 0, form.OptionInt("missing"))

	require.True(t, form.OptionBool("compact"))
	require.False(t, form.OptionBool("missing"))
}

func TestElement_NewElement(t *testing.T) {
	element := NewElement()
	require.NotNil(t, element.Options)
	require.NotNil(t, element.Children)
	require.Equal(t, 0, len(element.Children))
}

func TestElement_IsEmpty(t *testing.T) {
	require.True(t, Element{}.IsEmpty())
	require.False(t, Element{Type: "test"}.IsEmpty())
}

func TestElement_GetString(t *testing.T) {

	s := schema.New(schema.Object{Properties: schema.ElementMap{"name": schema.String{}}})
	element := Element{Type: "test", Path: "name"}
	value := mapof.Any{"name": "Sarah"}

	require.Equal(t, "Sarah", element.GetString(value, &s))

	// With no schema, getValue returns nil -> empty string
	require.Equal(t, "", element.GetString(value, nil))
}

func TestElement_GetSliceOfString(t *testing.T) {

	s := schema.New(schema.Object{
		Properties: schema.ElementMap{
			"tags": schema.Array{Items: schema.String{}},
		},
	})
	element := Element{Type: "test", Path: "tags"}
	value := mapof.Any{"tags": []string{"a", "b"}}

	require.Equal(t, []string{"a", "b"}, element.GetSliceOfString(value, &s))
}

func TestElement_GetSchema(t *testing.T) {

	s := schema.New(schema.Object{Properties: schema.ElementMap{"name": schema.String{}}})
	element := Element{Path: "name"}

	schemaElement := element.GetSchema(&s)
	require.NotNil(t, schemaElement)

	// With a nil schema, GetSchema returns nil
	require.Nil(t, element.GetSchema(nil))
}

func TestElement_EncodingAndWidget(t *testing.T) {
	useTestWidget()
	element := Element{Type: "test"}
	require.Equal(t, "", element.Encoding())

	// An unknown widget type produces an error from Widget()
	unknown := Element{Type: "does-not-exist"}
	_, err := unknown.Widget()
	require.Error(t, err)

	// Encoding of an unknown widget falls back to ""
	require.Equal(t, "", unknown.Encoding())
}

func TestElement_EditReadOnlyUsesView(t *testing.T) {
	useTestWidget()
	form, value := testForm(t)

	element := Element{Type: "test", Path: "name", ReadOnly: true}
	builder := html.New()
	require.NoError(t, element.Edit(&form, nil, value, builder))
	// A read-only element renders the View widget, not the Edit widget
	require.Equal(t, `<widget-view name="name">`, builder.String())
}

func TestElement_IsInputVisible(t *testing.T) {

	s := schema.New(schema.Object{Properties: schema.ElementMap{"name": schema.String{}}})

	// ReadOnly elements are never visible
	readOnly := Element{ReadOnly: true}
	visible, err := readOnly.isInputVisible(&s, nil)
	require.NoError(t, err)
	require.False(t, visible)

	// No show-if means always visible
	plain := Element{Path: "name"}
	visible, err = plain.isInputVisible(&s, mapof.Any{"name": "x"})
	require.NoError(t, err)
	require.True(t, visible)
}
