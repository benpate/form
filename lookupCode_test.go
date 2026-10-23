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
	require.Equal(t, "1.00", lookupCodes[0].Value)
	require.Equal(t, "One", lookupCodes[0].Label)
	require.Equal(t, "2.00", lookupCodes[1].Value)
	require.Equal(t, "Two", lookupCodes[1].Label)
	require.Equal(t, "3.00", lookupCodes[2].Value)
	require.Equal(t, "Three", lookupCodes[2].Label)
	require.False(t, writable)
}

func TestGetLookupCode_SliceOfString(t *testing.T) {

	element := Element{
		Type: "select",
		Options: mapof.Template{
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

// nilGroupProvider is a LookupProvider that always reports an unrecognized group.
type nilGroupProvider struct{}

func (nilGroupProvider) Group(string) LookupGroup { return nil }

func TestGetLookupCode_NilGroup(t *testing.T) {

	// A provider that returns a nil group must not panic; it falls through to the enum
	element := Element{
		Type: "select",
		Options: mapof.Template{
			"provider": "missing",
			"enum":     "one,two",
		},
	}

	lookupCodes, writable := GetLookupCodes(&element, nil, nilGroupProvider{})

	require.Len(t, lookupCodes, 2)
	require.Equal(t, "one", lookupCodes[0].Value)
	require.Equal(t, "two", lookupCodes[1].Value)
	require.False(t, writable)
}

func TestGetLookupCode_Provider(t *testing.T) {

	// A provider that returns a real group is used in preference to the enum
	element := Element{
		Type: "select",
		Options: mapof.Template{
			"provider": "colors",
			"enum":     "ignored",
		},
	}

	lookupCodes, writable := GetLookupCodes(&element, nil, staticProvider{
		"colors": NewReadOnlyLookupGroup(NewLookupCode("red"), NewLookupCode("green")),
	})

	require.Len(t, lookupCodes, 2)
	require.Equal(t, "red", lookupCodes[0].Value)
	require.Equal(t, "green", lookupCodes[1].Value)
	require.False(t, writable)
}

// staticProvider is a LookupProvider backed by a map of named groups.
type staticProvider map[string]LookupGroup

func (p staticProvider) Group(name string) LookupGroup { return p[name] }

func TestGetLookupCode_SliceOfLookupCodes(t *testing.T) {

	element := Element{
		Type: "select",
		Options: mapof.Template{
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

// TestGetLookupCode_SliceOfAny_Fractional confirms that fractional enum values keep every digit
// they need, and two decimals when two decimals are exact
func TestGetLookupCode_SliceOfAny_Fractional(t *testing.T) {

	element, err := Parse(`{"type":"select","path":"test","options":{"enum":[{"label":"Half","value":0.5},{"label":"Two and a quarter","value":2.25},{"label":"Eighth","value":0.125},{"label":"Negative","value":-1.5}]}}`)

	require.Nil(t, err)

	lookupCodes, writable := GetLookupCodes(&element, nil, nil)

	require.Len(t, lookupCodes, 4)
	require.Equal(t, "0.50", lookupCodes[0].Value)
	require.Equal(t, "Half", lookupCodes[0].Label)
	require.Equal(t, "2.25", lookupCodes[1].Value)
	require.Equal(t, "0.125", lookupCodes[2].Value)
	require.Equal(t, "-1.50", lookupCodes[3].Value)
	require.False(t, writable)
}

// TestGetLookupCode_SchemaNumber_Fractional confirms that a Number schema's fractional enum keeps
// every digit it needs
func TestGetLookupCode_SchemaNumber_Fractional(t *testing.T) {

	element := Element{Type: "select", Path: "test"}
	lookupCodes, writable := GetLookupCodes(&element, schema.Number{Enum: []float64{0.5, 0.125}}, nil)

	require.Len(t, lookupCodes, 2)
	require.Equal(t, "0.50", lookupCodes[0].Value)
	require.Equal(t, "0.125", lookupCodes[1].Value)
	require.False(t, writable)
}

// TestGetLookupCode_Fractional_RoundTrip confirms that every fractional lookup value posts back as
// the number it came from, and passes the schema's enum
func TestGetLookupCode_Fractional_RoundTrip(t *testing.T) {

	enum := []float64{0.5, 2.25, 0.125, 1.005, -1.5}
	s := schema.New(schema.Object{Properties: schema.ElementMap{"test": schema.Number{Enum: enum}}})

	lookupCodes, _ := GetLookupCodes(&Element{Type: "select", Path: "test"}, schema.Number{Enum: enum}, nil)
	require.Len(t, lookupCodes, len(enum))

	for index, lookupCode := range lookupCodes {
		object := mapof.Any{}
		require.NoError(t, s.Set(&object, "test", lookupCode.Value), "posted value %q", lookupCode.Value)
		require.Equal(t, enum[index], object["test"], "posted value %q", lookupCode.Value)
	}
}

// TestGetLookupCode_SliceOfNumbers confirms that an enum of bare numbers becomes one lookup code
// per number, with every digit each one needs
func TestGetLookupCode_SliceOfNumbers(t *testing.T) {

	element, err := Parse(`{"type":"select","path":"test","options":{"enum":[1,0.5,0.125]}}`)

	require.Nil(t, err)

	lookupCodes, writable := GetLookupCodes(&element, nil, nil)

	require.Len(t, lookupCodes, 3)
	require.Equal(t, LookupCode{Value: "1", Label: "1"}, lookupCodes[0])
	require.Equal(t, LookupCode{Value: "0.50", Label: "0.50"}, lookupCodes[1])
	require.Equal(t, LookupCode{Value: "0.125", Label: "0.125"}, lookupCodes[2])
	require.False(t, writable)
}
