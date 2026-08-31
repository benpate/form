package widget

import (
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/null"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

// toggleForm builds a single-toggle form over a map-backed object, which is the shape that
// exposes the absent-vs-false collapse: the map simply has no key until something writes one.
func toggleForm(t *testing.T, element schema.Element) form.Form {

	t.Helper()

	UseAll()

	return form.Form{
		Schema: schema.New(schema.Object{
			Properties: schema.ElementMap{
				"data": schema.Object{Properties: schema.ElementMap{"flag": element}},
			},
		}),
		Element: form.Element{
			Type:    "toggle",
			Path:    "data.flag",
			Options: mapof.Any{"true-text": "ON", "false-text": "OFF"},
		},
	}
}

// toggleObject is a map-backed object whose "data" property holds the toggle's value
type toggleObject struct {
	Data mapof.Any
}

// GetPointer implements schema.PointerGetter
func (object *toggleObject) GetPointer(name string) (any, bool) {
	if name == "data" {
		return &object.Data, true
	}
	return nil, false
}

// TestToggle_DefaultTrue_Absent verifies that a toggle whose schema declares a TRUE default
// renders in its ON position before anything has been saved.  Without this the toggle shows
// OFF, and the form's first save writes that phantom FALSE into storage.
func TestToggle_DefaultTrue_Absent(t *testing.T) {

	f := toggleForm(t, schema.Boolean{Default: null.NewBool(true)})
	object := &toggleObject{Data: mapof.NewAny()}

	result, err := f.Editor(object, nil)
	require.NoError(t, err)
	require.Contains(t, result, `value="true"`)
}

// TestToggle_DefaultTrue_StoredFalse verifies that a STORED false still wins over the
// default.  The default only fills in for a value that is genuinely absent.
func TestToggle_DefaultTrue_StoredFalse(t *testing.T) {

	f := toggleForm(t, schema.Boolean{Default: null.NewBool(true)})
	object := &toggleObject{Data: mapof.Any{"flag": false}}

	result, err := f.Editor(object, nil)
	require.NoError(t, err)
	require.NotContains(t, result, `value="true"`)
}

// TestToggle_NoDefault_Absent verifies that a toggle with no declared default is unchanged:
// an absent value still renders OFF.
func TestToggle_NoDefault_Absent(t *testing.T) {

	f := toggleForm(t, schema.Boolean{})
	object := &toggleObject{Data: mapof.NewAny()}

	result, err := f.Editor(object, nil)
	require.NoError(t, err)
	require.NotContains(t, result, `value="true"`)
}

// TestToggle_View_DefaultTrue verifies that the read-only view agrees with the editor, so a
// setting does not read as OFF on a profile page while its form shows ON.
func TestToggle_View_DefaultTrue(t *testing.T) {

	f := toggleForm(t, schema.Boolean{Default: null.NewBool(true)})
	object := &toggleObject{Data: mapof.NewAny()}

	result, err := f.Viewer(object, nil)
	require.NoError(t, err)
	require.Contains(t, result, "ON")
}
