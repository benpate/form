package widget

import (
	"regexp"
	"strings"
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

// iconLookupProvider returns LookupCodes that exercise every branch of iconName:
// a bare Value, an explicit Icon that overrides the Value, and a hostile name.
type iconLookupProvider struct{}

// Group returns the icon LookupCodes, regardless of the group name requested.
func (iconLookupProvider) Group(_ string) form.LookupGroup {
	return form.NewReadOnlyLookupGroup(
		form.LookupCode{Value: "folder", Label: "Folder", Group: "Navigation"},
		form.LookupCode{Value: "star", Label: "Star", Group: "Navigation"},
		form.LookupCode{Value: "REPLY", Label: "Reply", Icon: "reply-fill", Group: "Content"},
		form.LookupCode{Value: "bad name", Label: "Hostile", Group: "Content"},
	)
}

// selectIconsForm returns a Form that draws a select-icons widget over the "color" path.
func selectIconsForm(required bool) form.Form {

	UseAll()

	return form.New(
		schema.New(schema.Object{
			Properties: schema.ElementMap{
				"color": schema.String{Required: required},
			},
		}),
		form.Element{
			Type:    "select-icons",
			Path:    "color",
			Label:   "Icon",
			Options: map[string]any{"provider": "icons"},
		},
	)
}

// editSelectIcons renders the editor for the "color" path holding the given value.
func editSelectIcons(t *testing.T, required bool, value string) string {
	t.Helper()

	f := selectIconsForm(required)
	result, err := f.Editor(mapof.Any{"color": value}, iconLookupProvider{})
	require.NoError(t, err)

	return result
}

// isChecked reports whether the radio carrying this value is checked, without
// depending on the order attributes happen to be written in.
func isChecked(html string, value string) bool {
	pattern := regexp.MustCompile(`<input[^>]*value="` + regexp.QuoteMeta(value) + `"[^>]*checked="true"`)
	return pattern.MatchString(html)
}

func TestSelectIcons_ChecksTheMatchingTile(t *testing.T) {

	result := editSelectIcons(t, false, "star")

	require.True(t, isChecked(result, "star"))
	require.False(t, isChecked(result, "folder"))
}

func TestSelectIcons_PrefersIconOverValue(t *testing.T) {

	result := editSelectIcons(t, false, "")

	// The "REPLY" code carries Icon:"reply-fill", which must win over its Value
	require.Contains(t, result, "bi-reply-fill")
	require.NotContains(t, result, "bi-REPLY")
}

func TestSelectIcons_RejectsAnUnsafeIconName(t *testing.T) {

	result := editSelectIcons(t, false, "")

	// "bad name" would inject a second class, so no <i> is drawn for it at all
	require.NotContains(t, result, "bi-bad")
	require.Contains(t, result, "select-icons-blank")

	// ...but the tile itself is still offered, with its Value intact
	require.Contains(t, result, `value="bad name"`)
}

func TestSelectIcons_OffersNoneOnlyWhenNotRequired(t *testing.T) {

	optional := editSelectIcons(t, false, "")
	require.Contains(t, optional, `value=""`)
	require.Contains(t, optional, `aria-label="None"`)

	required := editSelectIcons(t, true, "")
	require.NotContains(t, required, `value=""`)
	require.NotContains(t, required, `aria-label="None"`)
}

func TestSelectIcons_RequiredAndUnmatchedChecksTheFirstTile(t *testing.T) {

	// Mirrors <select required> with no blank <option>: the first entry wins
	result := editSelectIcons(t, true, "")
	require.True(t, isChecked(result, "folder"))
}

func TestSelectIcons_OptionalAndUnmatchedChecksNone(t *testing.T) {

	result := editSelectIcons(t, false, "")

	require.True(t, isChecked(result, ""))
	require.False(t, isChecked(result, "folder"))
}

func TestSelectIcons_TilesCarryNoVisibleText(t *testing.T) {

	result := editSelectIcons(t, false, "star")

	// Tiles are icon-only, like the emoji picker. The name lives in title and
	// aria-label, which is the only accessible name a tile has.
	require.NotContains(t, result, "select-icons-tile-label")
	require.Contains(t, result, `title="Folder"`)
	require.Contains(t, result, `aria-label="Folder"`)

	// ...but the trigger button still shows the selected name as text
	require.Contains(t, result, `class="select-icons-button-label">Star<`)
}

func TestSelectIcons_DrawsGroupHeadersOncePerGroup(t *testing.T) {

	result := editSelectIcons(t, false, "")

	require.Equal(t, 1, strings.Count(result, `class="select-icons-header">Navigation<`))
	require.Equal(t, 1, strings.Count(result, `class="select-icons-header">Content<`))
}

func TestSelectIcons_TriggerButtonNeverSubmits(t *testing.T) {

	result := editSelectIcons(t, false, "star")

	// A bare <button> inside a <form> defaults to type="submit"
	require.Contains(t, result, `<button id="select-icons-color-button" type="button"`)
	require.Contains(t, result, `popovertarget="select-icons-color-popover"`)
}

func TestSelectIcons_PopoverAttributeIsWritten(t *testing.T) {

	result := editSelectIcons(t, false, "")

	// Attr() drops empty values, so the attribute has to carry "auto"
	require.Contains(t, result, `popover="auto"`)
}

func TestSelectIcons_TileIDsAreIndexed(t *testing.T) {

	// "bad name" as a DOM id would be invalid; tiles are keyed by index instead
	result := editSelectIcons(t, false, "")

	require.Contains(t, result, `id="select-icons-color-4"`)
	require.NotContains(t, result, `id="select-icons-color-bad name"`)
}

func TestSelectIcons_ViewShowsIconAndLabel(t *testing.T) {

	f := selectIconsForm(false)
	result, err := f.Viewer(mapof.Any{"color": "star"}, iconLookupProvider{})

	require.NoError(t, err)
	require.Contains(t, result, "bi-star")
	require.Contains(t, result, "Star")
}

func TestSelectIcons_ViewOfAnUnmatchedValueIsEmpty(t *testing.T) {

	f := selectIconsForm(false)
	result, err := f.Viewer(mapof.Any{"color": "nonexistent"}, iconLookupProvider{})

	require.NoError(t, err)
	require.NotContains(t, result, "nonexistent")
}
