package form

import (
	"testing"

	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

func TestLookupCodeSchema(t *testing.T) {

	test := func(name string, value string) {
		lookupCode := LookupCode{}
		s := schema.New(LookupCodeSchema())

		err := s.Set(&lookupCode, name, value)
		require.Nil(t, err)

		result, err := s.Get(&lookupCode, name)
		require.Nil(t, err)

		require.Equal(t, value, result)
	}

	test("value", "A VALUE")
	test("label", "A LABEL")
	test("description", "A DESCRIPTION")
	test("icon", "A ICON")
	test("group", "A GROUP")
	test("href", "http://href.com")
}

func TestGetLookupCode_String(t *testing.T) {

	element, err := Parse(`{"type":"select","path":"test","options":{"enum":"one,two,three"}}`)

	require.Nil(t, err)

	lookupCodes, writable := GetLookupCodes(&element, nil, nil)

	require.Len(t, lookupCodes, 3)
	require.Equal(t, "one", lookupCodes[0].Value)
	require.Equal(t, "one", lookupCodes[0].Label)
	require.Equal(t, "two", lookupCodes[1].Value)
	require.Equal(t, "two", lookupCodes[1].Label)
	require.Equal(t, "three", lookupCodes[2].Value)
	require.Equal(t, "three", lookupCodes[2].Label)
	require.False(t, writable)
}

func TestGetLookupCode_Slice(t *testing.T) {

	element, err := Parse(`{"type":"select","path":"test","options":{"enum":["one","two","three"]}}`)

	require.Nil(t, err)

	lookupCodes, writable := GetLookupCodes(&element, nil, nil)

	require.Len(t, lookupCodes, 3)
	require.Equal(t, "one", lookupCodes[0].Value)
	require.Equal(t, "one", lookupCodes[0].Label)
	require.Equal(t, "two", lookupCodes[1].Value)
	require.Equal(t, "two", lookupCodes[1].Label)
	require.Equal(t, "three", lookupCodes[2].Value)
	require.Equal(t, "three", lookupCodes[2].Label)
	require.False(t, writable)
}

func TestGetLookupCode_SliceOfAny(t *testing.T) {

	element, err := Parse(`{"type":"select","path":"test","options":{"enum":[{"label":"One","value":1},{"label":"Two","value":2},{"label":"Three","value":3}]}}`)

	require.Nil(t, err)

	lookupCodes, writable := GetLookupCodes(&element, nil, nil)

	require.Len(t, lookupCodes, 3)
	require.Equal(t, "1", lookupCodes[0].Value)
	require.Equal(t, "One", lookupCodes[0].Label)
	require.Equal(t, "2", lookupCodes[1].Value)
	require.Equal(t, "Two", lookupCodes[1].Label)
	require.Equal(t, "3", lookupCodes[2].Value)
	require.Equal(t, "Three", lookupCodes[2].Label)
	require.False(t, writable)
}

func TestGetLookupCode_SliceOfString(t *testing.T) {

	element := Element{
		Type: "select",
		Options: mapof.Any{
			"enum": []string{"one", "two", "three"},
		},
	}

	lookupCodes, writable := GetLookupCodes(&element, nil, nil)

	require.Len(t, lookupCodes, 3)
	require.Equal(t, "one", lookupCodes[0].Value)
	require.Equal(t, "one", lookupCodes[0].Label)
	require.Equal(t, "two", lookupCodes[1].Value)
	require.Equal(t, "two", lookupCodes[1].Label)
	require.Equal(t, "three", lookupCodes[2].Value)
	require.Equal(t, "three", lookupCodes[2].Label)
	require.False(t, writable)
}

func TestGetLookupCode_SliceOfLookupCodes(t *testing.T) {

	element := Element{
		Type: "select",
		Options: mapof.Any{
			"enum": []LookupCode{
				{Value: "1", Label: "One"},
				{Value: "2", Label: "Two"},
				{Value: "3", Label: "Three"},
			},
		},
	}

	lookupCodes, writable := GetLookupCodes(&element, nil, nil)

	require.Len(t, lookupCodes, 3)
	require.Equal(t, "1", lookupCodes[0].Value)
	require.Equal(t, "One", lookupCodes[0].Label)
	require.Equal(t, "2", lookupCodes[1].Value)
	require.Equal(t, "Two", lookupCodes[1].Label)
	require.Equal(t, "3", lookupCodes[2].Value)
	require.Equal(t, "Three", lookupCodes[2].Label)
	require.False(t, writable)
}
