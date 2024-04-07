package form

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestElement_AllElements(t *testing.T) {

	useTestWidget()

	element := Element{
		Type: "test",
		Children: []Element{{
			Type: "test",
			Path: "name",
		}, {
			Type: "test",
			Path: "email",
		}, {
			Type: "test",
			Path: "phone",
		}},
	}

	elements := element.AllElements()

	require.Equal(t, 3, len(elements))
	require.Equal(t, "test", elements[0].Type)
	require.Equal(t, "name", elements[0].Path)
	require.Equal(t, "test", elements[1].Type)
	require.Equal(t, "email", elements[1].Path)
	require.Equal(t, "test", elements[2].Type)
	require.Equal(t, "phone", elements[2].Path)
}
