package form

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalMap(t *testing.T) {

	elementMap := map[string]any{
		"type":     "text",
		"path":     "name",
		"label":    "Name",
		"readOnly": true,
	}

	element := MustParse(elementMap)

	require.Equal(t, "text", element.Type)
	require.Equal(t, "name", element.Path)
	require.Equal(t, "Name", element.Label)
	require.True(t, element.ReadOnly)
}
