package form

import (
	"testing"

	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
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

// TestElement_GetString_Fractional confirms that a stored fraction renders with every digit it needs,
// and two decimals when two decimals are exact
func TestElement_GetString_Fractional(t *testing.T) {

	s := schema.New(schema.Object{Properties: schema.ElementMap{"amount": schema.Number{}}})
	element := Element{Type: "text", Path: "amount"}

	require.Equal(t, "1.50", element.GetString(&mapof.Any{"amount": 1.5}, &s))
	require.Equal(t, "0.125", element.GetString(&mapof.Any{"amount": 0.125}, &s))
	require.Equal(t, "40.712776", element.GetString(&mapof.Any{"amount": 40.712776}, &s))
}

// TestElement_GetSliceOfString_Fractional confirms that every stored fraction in a list renders with
// every digit it needs
func TestElement_GetSliceOfString_Fractional(t *testing.T) {

	s := schema.New(schema.Object{Properties: schema.ElementMap{"amounts": schema.Array{Items: schema.Number{}}}})
	element := Element{Type: "multiselect", Path: "amounts"}

	require.Equal(t, []string{"1.50", "0.125"}, element.GetSliceOfString(&mapof.Any{"amounts": []any{1.5, 0.125}}, &s))
	require.Equal(t, []string{"one", "two"}, element.GetSliceOfString(&mapof.Any{"amounts": []any{"one", "two"}}, &s))
}
