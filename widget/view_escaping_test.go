package widget

import (
	"strings"
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

// hostileLookupProvider returns LookupCodes whose Labels and Groups contain markup,
// standing in for a provider backed by text that an end-user typed.
type hostileLookupProvider struct{}

// Group returns the hostile LookupCodes, regardless of the group name requested.
func (hostileLookupProvider) Group(_ string) form.LookupGroup {
	return form.NewReadOnlyLookupGroup(
		form.LookupCode{Value: "ONE", Label: `<script>alert(1)</script>`, Group: `<b>GRP</b>`},
		form.LookupCode{Value: "TWO", Label: `Salt & Pepper`, Group: `<b>GRP</b>`},
	)
}

// viewEscapingWidgets lists every widget whose View draws LookupCode text.
var viewEscapingWidgets = []string{
	"checkbox",
	"check-button-group",
	"multiselect",
	"radio",
	"radio-button-group",
	"radio-button-group-horizontal",
	"radio-colors",
	"select",
	"select-group",
	"select-icons",
}

// hostileForm returns a Form that draws the named widget from the hostile provider.
func hostileForm(widgetType string) form.Form {

	UseAll()

	return form.New(
		schema.New(schema.Object{
			Properties: schema.ElementMap{
				"tags": schema.Array{Items: schema.String{}},
			},
		}),
		form.Element{
			Type:    widgetType,
			Path:    "tags",
			Options: mapof.Any{"provider": "test"},
		},
	)
}

func TestView_LookupCodesAreEscaped(t *testing.T) {

	for _, widgetType := range viewEscapingWidgets {

		f := hostileForm(widgetType)
		result, err := f.Viewer(mapof.Any{"tags": []string{"ONE", "TWO"}}, hostileLookupProvider{})

		require.Nil(t, err)
		require.NotContains(t, result, "<script>", widgetType)
		require.NotContains(t, result, "<b>", widgetType)
	}
}

func TestView_ValueDivIsWellFormed(t *testing.T) {

	for _, widgetType := range viewEscapingWidgets {

		f := hostileForm(widgetType)
		result, err := f.Viewer(mapof.Any{"tags": []string{"ONE", "TWO"}}, hostileLookupProvider{})

		require.Nil(t, err)

		// The wrapper's end bracket must be written before any of its text, or the
		// text is parsed as part of the opening tag instead of as content.
		require.True(t, strings.HasPrefix(result, `<div class="layout-value">`), widgetType+": "+result)
		require.True(t, strings.HasSuffix(result, `</div>`), widgetType+": "+result)
	}
}

func TestTextArea_RowsOmittedWhenUnset(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{Type: "textarea", Path: "username"},
	)

	result, err := f.Editor(nil, testLookupProvider{})

	require.Nil(t, err)
	require.NotContains(t, result, "rows=")
}

func TestHeading_EditMatchesView(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{Type: "heading", Label: `Salt & <b>Pepper</b>`},
	)

	viewer, err := f.Viewer(nil, nil)
	require.Nil(t, err)

	editor, err := f.Editor(nil, nil)
	require.Nil(t, err)

	require.Equal(t, viewer, editor)
	require.NotContains(t, editor, "<b>")
}
