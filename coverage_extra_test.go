package form

import (
	"testing"

	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

func TestGetSchemaEnumeration_AllTypes(t *testing.T) {

	// String enum
	{
		codes, _ := GetLookupCodes(&Element{}, schema.String{Enum: []string{"a", "b"}}, nil)
		require.Equal(t, 2, len(codes))
	}
	// Array of strings
	{
		codes, _ := GetLookupCodes(&Element{}, schema.Array{Items: schema.String{Enum: []string{"x"}}}, nil)
		require.Equal(t, 1, len(codes))
		require.Equal(t, "x", codes[0].Value)
	}
	// Integer enum
	{
		codes, _ := GetLookupCodes(&Element{}, schema.Integer{Enum: []int{1, 2, 3}}, nil)
		require.Equal(t, 3, len(codes))
	}
	// Number enum
	{
		codes, _ := GetLookupCodes(&Element{}, schema.Number{Enum: []float64{1.5, 2.5}}, nil)
		require.Equal(t, 2, len(codes))
	}
	// A type with no enumeration yields no codes
	{
		codes, _ := GetLookupCodes(&Element{}, schema.Boolean{}, nil)
		require.Equal(t, 0, len(codes))
	}
}

func TestGetLookupCodes_FromProvider(t *testing.T) {

	group := NewReadOnlyLookupGroup(NewLookupCode("a"), NewLookupCode("b"))
	provider := mockProvider{group: group}

	element := &Element{Options: map[string]any{"provider": "things"}}

	codes, writable := GetLookupCodes(element, nil, provider)
	require.False(t, writable) // ReadOnlyLookupGroup is not writable
	require.Equal(t, 2, len(codes))
}

func TestGetLookupCodes_EnumSliceTypes(t *testing.T) {

	// []string enum
	{
		element := &Element{Options: map[string]any{"enum": []string{"a", "b"}}}
		codes, _ := GetLookupCodes(element, nil, nil)
		require.Equal(t, 2, len(codes))
	}
	// []LookupCode enum
	{
		element := &Element{Options: map[string]any{"enum": []LookupCode{{Value: "x"}}}}
		codes, _ := GetLookupCodes(element, nil, nil)
		require.Equal(t, 1, len(codes))
	}
	// []any enum
	{
		element := &Element{Options: map[string]any{"enum": []any{"a", "b", "c"}}}
		codes, _ := GetLookupCodes(element, nil, nil)
		require.Equal(t, 3, len(codes))
	}
}

func TestSortLookupCodeByGroupThenLabel(t *testing.T) {

	// Different groups sort by group
	require.Equal(t, -1, SortLookupCodeByGroupThenLabel(
		LookupCode{Group: "a"}, LookupCode{Group: "b"}))
	require.Equal(t, 1, SortLookupCodeByGroupThenLabel(
		LookupCode{Group: "b"}, LookupCode{Group: "a"}))

	// Same group sorts by label
	require.Equal(t, -1, SortLookupCodeByGroupThenLabel(
		LookupCode{Group: "g", Label: "a"}, LookupCode{Group: "g", Label: "b"}))
	require.Equal(t, 1, SortLookupCodeByGroupThenLabel(
		LookupCode{Group: "g", Label: "b"}, LookupCode{Group: "g", Label: "a"}))

	// Identical group and label are equal
	require.Equal(t, 0, SortLookupCodeByGroupThenLabel(
		LookupCode{Group: "g", Label: "x"}, LookupCode{Group: "g", Label: "x"}))
}

func TestLookupCode_GetStringOK(t *testing.T) {

	code := LookupCode{Value: "v", Label: "l", Description: "d", Icon: "i", Group: "g", Href: "h"}

	for name, expected := range map[string]string{
		"value": "v", "label": "l", "description": "d",
		"icon": "i", "group": "g", "href": "h",
	} {
		result, ok := code.GetStringOK(name)
		require.True(t, ok, name)
		require.Equal(t, expected, result, name)
	}

	_, ok := code.GetStringOK("unknown")
	require.False(t, ok)
}

func TestLookupCode_SetString(t *testing.T) {

	code := &LookupCode{}

	for _, name := range []string{"value", "label", "description", "icon", "group", "href"} {
		require.True(t, code.SetString(name, "x"), name)
	}

	require.Equal(t, "x", code.Value)
	require.Equal(t, "x", code.Href)

	// Unknown property is rejected
	require.False(t, code.SetString("unknown", "x"))
}

func TestElement_UnmarshalMap_WithChildren(t *testing.T) {

	element := Element{}
	err := element.UnmarshalMap(map[string]any{
		"type":     "layout",
		"id":       "root",
		"label":    "Root",
		"readOnly": true,
		"options":  map[string]any{"key": "value"},
		"children": []any{
			map[string]any{"type": "text", "path": "name"},
			map[string]any{"type": "text", "path": "email"},
		},
	})

	require.NoError(t, err)
	require.Equal(t, "layout", element.Type)
	require.Equal(t, "root", element.ID)
	require.True(t, element.ReadOnly)
	require.Equal(t, "value", element.Options.GetString("key"))
	require.Equal(t, 2, len(element.Children))
	require.Equal(t, "name", element.Children[0].Path)
}

func TestElement_UnmarshalMap_InvalidChild(t *testing.T) {

	// A child that is not a map produces an error
	element := Element{}
	err := element.UnmarshalMap(map[string]any{
		"children": []any{"not a map"},
	})
	require.Error(t, err)
}

func TestElement_AllElements_ReadOnlyAndRoot(t *testing.T) {

	// A read-only root yields no elements
	readOnly := Element{ReadOnly: true, Path: "x"}
	require.Equal(t, 0, len(readOnly.AllElements()))

	// A root with no path yields only its children's paths
	root := Element{
		Children: []Element{
			{Path: "a"},
			{Path: "b"},
		},
	}
	require.Equal(t, 2, len(root.AllElements()))
}
