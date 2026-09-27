package widget

import (
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/stretchr/testify/require"
)

// TestOptionTemplates_RenderPerValue parses one form and renders it for two values, and
// requires that each render carries its own value while the parsed form is unchanged
func TestOptionTemplates_RenderPerValue(t *testing.T) {

	// Emissary BUG-204: rendered options were written back into the shared form, so the first
	// render after a restart fixed every later form to its value.
	UseAll()
	element := parseOptionTemplateForm(t)

	first, err := form.Editor(getTestSchema(), element, mapof.Any{"username": "first", "name": "a.png"}, nil)
	require.NoError(t, err)

	second, err := form.Editor(getTestSchema(), element, mapof.Any{"username": "second", "name": "b.png"}, nil)
	require.NoError(t, err)

	require.Contains(t, first, `install validator(url:&#39;/v?user=first&#39;)`)
	require.Contains(t, first, `hx-post="/first/delete"`)
	require.Contains(t, second, `install validator(url:&#39;/v?user=second&#39;)`)
	require.Contains(t, second, `hx-post="/second/delete"`)
	require.NotContains(t, second, "first")

	// The parsed form still holds its templates
	require.Equal(t, "/v?user={{.username}}", element.Children[0].Options.GetString("validator", nil))
	require.Equal(t, "/{{.username}}/delete", element.Children[1].Options.GetString("delete", nil))
}

// TestOptionTemplates_AnyOption requires that options of every type render against the value, not
// only the URL options: a placeholder string, a "rows" number, and a "required" bool
func TestOptionTemplates_AnyOption(t *testing.T) {

	UseAll()

	element, err := form.Parse(map[string]any{
		"type": "layout-vertical",
		"children": []any{
			map[string]any{"type": "text", "path": "name", "options": map[string]any{
				"placeholder": "Hi {{.username}}",
				"required":    "{{if .locked}}true{{else}}false{{end}}",
			}},
			map[string]any{"type": "textarea", "path": "name", "options": map[string]any{"rows": "{{.rows}}"}},
		},
	})
	require.NoError(t, err)

	locked, err := form.Editor(getTestSchema(), element, mapof.Any{"username": "first", "locked": true, "rows": 7}, nil)
	require.NoError(t, err)
	require.Contains(t, locked, `placeholder="Hi first"`)
	require.Contains(t, locked, `required="true"`)
	require.Contains(t, locked, `rows="7"`)

	open, err := form.Editor(getTestSchema(), element, mapof.Any{"username": "second", "locked": false, "rows": 3}, nil)
	require.NoError(t, err)
	require.Contains(t, open, `placeholder="Hi second"`)
	require.NotContains(t, open, `required="true"`)
	require.Contains(t, open, `rows="3"`)
}

// TestOptionTemplates_Concurrent renders one parsed form from many goroutines, which the race
// detector reports if any render writes the shared form
func TestOptionTemplates_Concurrent(t *testing.T) {

	UseAll()
	element := parseOptionTemplateForm(t)
	mismatches := make(chan string, 20)

	var wait sync.WaitGroup
	for index := range 20 {
		wait.Go(func() {
			username := "user" + strconv.Itoa(index)
			result, err := form.Editor(getTestSchema(), element, mapof.Any{"username": username, "name": "a.png"}, nil)

			// Report a mismatch rather than failing here, because require cannot run off the test goroutine
			if (err != nil) || !strings.Contains(result, `url:&#39;/v?user=`+username+`&#39;`) {
				mismatches <- username
			}
		})
	}

	wait.Wait()
	close(mismatches)

	for username := range mismatches {
		t.Errorf("render for %q did not carry its own value", username)
	}
}

// TestOptionTemplates_Errors requires that a template that fails to render fails the form, in
// each widget that renders one
func TestOptionTemplates_Errors(t *testing.T) {

	UseAll()
	value := mapof.Any{"username": "first", "name": "a.png"}

	for _, element := range []form.Element{
		{Type: "text", Path: "username", Options: mapof.Template{"validator": "/v?user={{.username.missing}}"}},
		{Type: "upload", Path: "name", Options: mapof.Template{"delete": "/{{.username.missing}}/delete"}},
		{Type: "html-remote", Options: mapof.Template{"url": "/{{.username.missing}}"}},
	} {
		_, err := form.Editor(getTestSchema(), element, value, nil)
		require.Error(t, err, element.Type)
	}
}

// parseOptionTemplateForm parses a form whose text and upload elements carry option templates,
// as a Template's hjson does
func parseOptionTemplateForm(t *testing.T) form.Element {

	t.Helper()

	element, err := form.Parse(map[string]any{
		"type": "layout-vertical",
		"children": []any{
			map[string]any{"type": "text", "path": "username", "options": map[string]any{"validator": "/v?user={{.username}}"}},
			map[string]any{"type": "upload", "path": "name", "options": map[string]any{"delete": "/{{.username}}/delete"}},
		},
	})

	require.NoError(t, err)
	return element
}
