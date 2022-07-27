package form

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestElement_AllElements(t *testing.T) {

	element := Element{
		Type: "layout-vertical",
		Children: []Element{{
			Type: "textarea",
			Path: "name",
		}, {
			Type: "text",
			Path: "email",
		}, {
			Type: "number",
			Path: "phone",
		}},
	}

	elements := element.AllElements()

	require.Equal(t, 3, len(elements))
	require.Equal(t, "textarea", elements[0].Type)
	require.Equal(t, "name", elements[0].Path)
	require.Equal(t, "text", elements[1].Type)
	require.Equal(t, "email", elements[1].Path)
	require.Equal(t, "number", elements[2].Type)
	require.Equal(t, "phone", elements[2].Path)
}
