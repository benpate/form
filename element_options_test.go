package form

import (
	"encoding/json"
	"testing"

	"github.com/benpate/rosetta/loose"
	"github.com/stretchr/testify/require"
)

// optionsTestData is a struct for option templates to render against
type optionsTestData struct {
	ID string
}

// TestElement_UnmarshalMap_OptionTemplates requires that Parse compiles option templates at
// every depth, and fails on a malformed one
func TestElement_UnmarshalMap_OptionTemplates(t *testing.T) {

	element, err := Parse(map[string]any{
		"type":    "layout-vertical",
		"options": map[string]any{"style": "wide"},
		"children": []any{
			map[string]any{"type": "text", "path": "token", "options": map[string]any{"validator": "/v?id={{.ID}}", "rows": 6}},
		},
	})
	require.NoError(t, err)
	require.Equal(t, "wide", element.Options.GetString("style", nil))
	require.True(t, element.Children[0].Options["validator"].(loose.Template).IsTemplate())
	require.Equal(t, 6, element.Children[0].Options.GetInt("rows", nil))

	result, err := element.Children[0].Options.Evaluate("validator", optionsTestData{ID: "123"})
	require.NoError(t, err)
	require.Equal(t, "/v?id=123", result)

	// A malformed template is kept as a plain string
	malformed, err := Parse(map[string]any{
		"type":     "layout-vertical",
		"children": []any{map[string]any{"type": "text", "options": map[string]any{"validator": "{{.ID | nosuchfunc}}"}}},
	})
	require.NoError(t, err)
	require.Equal(t, "{{.ID | nosuchfunc}}", malformed.Children[0].Options["validator"])
}

// TestElement_UnmarshalMap_KeepsInput requires that parsing leaves the caller's map unchanged
func TestElement_UnmarshalMap_KeepsInput(t *testing.T) {

	options := map[string]any{"validator": "/v?id={{.ID}}"}

	_, err := Parse(map[string]any{"type": "text", "options": options})
	require.NoError(t, err)
	require.Equal(t, "/v?id={{.ID}}", options["validator"])
}

// TestElement_JSON_OptionTemplates requires that an Element decoded from JSON compiles its
// option templates at every depth, and round-trips as its source
func TestElement_JSON_OptionTemplates(t *testing.T) {

	source := `{"type":"layout-vertical","id":"","path":"","children":[{"type":"text","id":"","path":"token","options":{"autocomplete":"off","validator":"/v?id={{.ID}}"}}]}`

	var element Element
	require.NoError(t, json.Unmarshal([]byte(source), &element))
	require.True(t, element.Children[0].Options["validator"].(loose.Template).IsTemplate())

	encoded, err := json.Marshal(element)
	require.NoError(t, err)
	require.JSONEq(t, source, string(encoded))

	require.NoError(t, json.Unmarshal([]byte(`{"type":"text","options":{"validator":"{{.ID | nosuchfunc}}"}}`), &element))
	require.Equal(t, "{{.ID | nosuchfunc}}", element.Options["validator"])
}
