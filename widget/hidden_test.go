package widget

import (
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/stretchr/testify/require"
)

// TestHidden_OptionValue requires that a "value" option, parsed or written in Go, replaces the
// value from the object
func TestHidden_OptionValue(t *testing.T) {

	UseAll()

	parsed, err := form.Parse(map[string]any{"type": "hidden", "path": "name", "options": map[string]any{"value": "from-option"}})
	require.NoError(t, err)

	for _, element := range []form.Element{
		parsed,
		{Type: "hidden", Path: "name", Options: mapof.Template{"value": "from-option"}},
	} {
		result, err := form.Editor(getTestSchema(), element, mapof.Any{"name": "from-object"}, nil)
		require.NoError(t, err)
		require.Contains(t, result, `value="from-option"`)
		require.NotContains(t, result, "from-object")
	}

	// Without the option, the object's value is used
	result, err := form.Editor(getTestSchema(), form.Element{Type: "hidden", Path: "name"}, mapof.Any{"name": "from-object"}, nil)
	require.NoError(t, err)
	require.Contains(t, result, `value="from-object"`)
}
