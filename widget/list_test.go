package widget

import (
	"net/url"
	"strings"
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/sliceof"
	"github.com/stretchr/testify/require"
)

// listForm returns a Form that draws a list widget over "tags", an Array of at most three Strings.
func listForm(options mapof.Template) form.Form {

	UseAll()

	return form.New(
		schema.New(schema.Object{
			Properties: schema.ElementMap{
				"tags":   schema.Array{MinLength: 1, MaxLength: 3, Items: schema.String{MaxLength: 20, Pattern: "^[a-z]+$"}},
				"name":   schema.String{},
				"counts": schema.Array{Items: schema.Integer{}},
			},
		}),
		form.Element{Type: "list", Path: "tags", Label: "Tags", Options: options},
	)
}

// listRowCount returns the number of editable rows drawn in the table, not counting the template's.
func listRowCount(html string) int {
	table, _, _ := strings.Cut(html, "<template")
	return strings.Count(table, `<tr class="grid-row">`)
}

func TestList_Edit(t *testing.T) {

	f := listForm(nil)
	value := mapof.Any{"tags": sliceof.String{"one", "two"}}

	result, err := f.Editor(&value, nil)
	require.NoError(t, err)

	row := func(value string, disabled string) string {
		return `<tr class="grid-row"><td class="grid-cell list-handle" aria-hidden="true"><i class="bi bi-grip-vertical"></i></td>` +
			`<td class="grid-cell"><input type="text" name="tags" value="` + value + `" aria-label="Tags" aria-describedby="list-tags-hint" autocomplete="off" maxlength="20" pattern="^[a-z]+$"></td>` +
			`<td class="grid-cell grid-controls"><button type="button" class="list-remove" aria-label="Remove this row"` + disabled + `><i class="bi bi-x-lg" aria-hidden="true"></i></button></td></tr>`
	}

	expected := `<div id="list-tags" class="list-widget" data-script="install ListWidget" data-minimum="1" data-maximum="3">` +
		`<table class="grid list"><tbody>` + row("one", "") + row("two", "") + `</tbody></table>` +
		`<template class="list-template">` + row("", "") + `</template>` +
		`<button type="button" class="link list-add"><i class="bi bi-plus-lg" aria-hidden="true"></i><span> Add a Row</span></button>` +
		`<div id="list-tags-hint" class="list-hint hide-visual">Enter adds a row below. Alt+Up and Alt+Down move this row, and Alt+Home and Alt+End move it to the top or bottom.</div>` +
		`<div class="list-status hide-visual" aria-live="polite"></div></div>`

	require.Equal(t, expected, result)
}

func TestList_Edit_EmptyDrawsOneBlankRow(t *testing.T) {

	f := listForm(nil)
	value := mapof.Any{}

	result, err := f.Editor(&value, nil)
	require.NoError(t, err)
	require.Equal(t, 1, listRowCount(result))
	require.Contains(t, result, `name="tags" value=""`)

	// RULE: The one blank row is also the minimum, so its X is disabled
	table, _, _ := strings.Cut(result, "<template")
	require.Contains(t, table, `aria-label="Remove this row" disabled="true"`)
}

func TestList_Edit_Limits(t *testing.T) {

	f := listForm(nil)

	// At the maximum, the add button is hidden
	full := mapof.Any{"tags": sliceof.String{"a", "b", "c"}}
	result, err := f.Editor(&full, nil)
	require.NoError(t, err)
	require.Equal(t, 3, listRowCount(result))
	require.Contains(t, result, `<button type="button" class="link list-add" hidden="true">`)

	// At the minimum, every X in the table is disabled, but the template's is not
	minimum := mapof.Any{"tags": sliceof.String{"a"}}
	result, err = f.Editor(&minimum, nil)
	require.NoError(t, err)

	table, template, _ := strings.Cut(result, "<template")
	require.Contains(t, table, `disabled="true"`)
	require.NotContains(t, template, `disabled="true"`)
	require.NotContains(t, result, `list-add" hidden`)
}

func TestList_Edit_AddLabel(t *testing.T) {

	f := listForm(mapof.Template{"add-label": "Add a Hashtag"})
	value := mapof.Any{}

	result, err := f.Editor(&value, nil)
	require.NoError(t, err)
	require.Contains(t, result, `<span> Add a Hashtag</span>`)
}

func TestList_Edit_Required(t *testing.T) {

	UseAll()

	f := form.New(
		schema.New(schema.Object{Properties: schema.ElementMap{
			"tags": schema.Array{Required: true, Items: schema.String{}},
		}}),
		form.Element{Type: "list", Path: "tags"},
	)

	value := mapof.Any{}
	result, err := f.Editor(&value, nil)
	require.NoError(t, err)
	require.Contains(t, result, `data-required="true"`)

	// RULE: A row never carries "required", because a blank row is dropped at Save
	require.NotContains(t, result, ` required=`)
}

func TestList_ValuesAreEscaped(t *testing.T) {

	f := listForm(nil)
	value := mapof.Any{"tags": sliceof.String{`"><script>alert(1)</script>`, `Salt & Pepper`}}

	editor, err := f.Editor(&value, nil)
	require.NoError(t, err)
	require.NotContains(t, editor, "<script>")
	require.Contains(t, editor, `value="&#34;&gt;&lt;script&gt;alert(1)&lt;/script&gt;"`)
	require.Contains(t, editor, `value="Salt &amp; Pepper"`)

	viewer, err := f.Viewer(&value, nil)
	require.NoError(t, err)
	require.NotContains(t, viewer, "<script>")
	require.Contains(t, viewer, `<td class="grid-cell">Salt &amp; Pepper</td>`)
}

func TestList_View(t *testing.T) {

	f := listForm(nil)

	value := mapof.Any{"tags": sliceof.String{"one", "two"}}
	result, err := f.Viewer(&value, nil)
	require.NoError(t, err)
	require.Equal(t, `<table class="grid list"><tbody><tr class="grid-row"><td class="grid-cell">one</td></tr><tr class="grid-row"><td class="grid-cell">two</td></tr></tbody></table>`, result)

	empty := mapof.Any{}
	result, err = f.Viewer(&empty, nil)
	require.NoError(t, err)
	require.Equal(t, `<div class="layout-value"></div>`, result)
}

func TestList_ReadOnlyEditDrawsTheView(t *testing.T) {

	f := listForm(nil)
	f.Element.ReadOnly = true
	value := mapof.Any{"tags": sliceof.String{"one"}}

	result, err := f.Editor(&value, nil)
	require.NoError(t, err)
	require.NotContains(t, result, "<input")
	require.Contains(t, result, `<td class="grid-cell">one</td>`)
}

func TestList_RefusesOtherSchemas(t *testing.T) {

	for _, path := range []string{"name", "counts", "missing"} {

		f := listForm(nil)
		f.Element.Path = path
		value := mapof.Any{"name": "John", "counts": sliceof.Int{1, 2}}

		_, err := f.Editor(&value, nil)
		require.Error(t, err, path)

		_, err = f.Viewer(&value, nil)
		require.Error(t, err, path)

		err = List{}.SetURLValue(&f, &f.Element, &value, url.Values{path: {"x"}})
		require.Error(t, err, path)
	}
}

func TestList_SetURLValues(t *testing.T) {

	tests := []struct {
		name     string
		posted   []string
		expected sliceof.String
	}{
		{name: "keeps the posted order", posted: []string{"c", "a", "b"}, expected: sliceof.String{"c", "a", "b"}},
		{name: "trims each value", posted: []string{"  a ", "\tb\n"}, expected: sliceof.String{"a", "b"}},
		{name: "drops blank rows", posted: []string{"", "a", "   ", "b", ""}, expected: sliceof.String{"a", "b"}},
		{name: "keeps the first of a repeat, ignoring case", posted: []string{"b", "a", "B", " a "}, expected: sliceof.String{"b", "a"}},
		{name: "keeps a later repeat out of the count", posted: []string{"a", "A", "a", "b", "c"}, expected: sliceof.String{"a", "b", "c"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			f := listForm(nil)
			f.Schema = schema.New(schema.Object{Properties: schema.ElementMap{
				"tags": schema.Array{MaxLength: 3, Items: schema.String{}},
			}})

			value := mapof.Any{"tags": sliceof.String{"old"}}
			require.NoError(t, f.SetURLValues(&value, url.Values{"tags": test.posted}, nil))

			result := f.Element.GetSliceOfString(value, &f.Schema)
			require.Equal(t, []string(test.expected), result)
		})
	}
}

func TestList_SetURLValues_EmptyClears(t *testing.T) {

	f := listForm(nil)
	f.Schema = schema.New(schema.Object{Properties: schema.ElementMap{
		"tags": schema.Array{Items: schema.String{}},
	}})

	// No rows posted, or only blank ones, clears the list
	for _, posted := range []url.Values{{}, {"tags": {"", " "}}} {
		value := mapof.Any{"tags": sliceof.String{"old"}}
		require.NoError(t, f.SetURLValues(&value, posted, nil))
		require.Empty(t, f.Element.GetSliceOfString(value, &f.Schema))
	}
}

func TestList_SetURLValues_RefusedListKeepsTheOldValue(t *testing.T) {

	f := listForm(nil)

	// Four distinct values exceed the maximum of three, so the schema refuses the whole list.
	// SetURLValues logs that at debug level and goes on, which is why the browser must stop it first.
	value := mapof.Any{"tags": sliceof.String{"old"}}
	require.NoError(t, f.SetURLValues(&value, url.Values{"tags": {"a", "b", "c", "d"}}, nil))
	require.Equal(t, []string{"old"}, f.Element.GetSliceOfString(value, &f.Schema))
}

func TestList_SetURLValues_ReadOnlyIsNotWritten(t *testing.T) {

	f := listForm(nil)
	f.Element.ReadOnly = true

	value := mapof.Any{"tags": sliceof.String{"old"}}
	require.NoError(t, f.SetURLValues(&value, url.Values{"tags": {"new"}}, nil))
	require.Equal(t, []string{"old"}, f.Element.GetSliceOfString(value, &f.Schema))
}

func TestCleanList(t *testing.T) {
	require.Equal(t, sliceof.String{}, cleanList(nil))
	require.Equal(t, sliceof.String{"Ä"}, cleanList([]string{"Ä", "ä"}))
	require.Equal(t, sliceof.String{"one two"}, cleanList([]string{" one two "}))
}
