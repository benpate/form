package form

import (
	"testing"

	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

func TestReadOnlyLookupGroup(t *testing.T) {

	group := NewReadOnlyLookupGroup(
		NewLookupCode("a"),
		LookupCode{Value: "b", Label: "Bravo"},
	)

	codes := group.Get()
	require.Equal(t, 2, len(codes))

	// Value returns the matching code
	require.Equal(t, "Bravo", group.Value("b").Label)

	// A missing value returns an empty LookupCode
	require.Equal(t, LookupCode{}, group.Value("missing"))
}

func TestFilterLookupCodeByGroup(t *testing.T) {

	filter := FilterLookupCodeByGroup("fruits")

	require.True(t, filter(LookupCode{Value: "apple", Group: "fruits"}))
	require.False(t, filter(LookupCode{Value: "carrot", Group: "vegetables"}))
}

func TestSortLookupCodeByLabel_Equal(t *testing.T) {

	// Equal labels return 0
	require.Equal(t, 0, SortLookupCodeByLabel(
		LookupCode{Label: "same"},
		LookupCode{Label: "same"},
	))
	require.Equal(t, -1, SortLookupCodeByLabel(LookupCode{Label: "a"}, LookupCode{Label: "b"}))
	require.Equal(t, 1, SortLookupCodeByLabel(LookupCode{Label: "b"}, LookupCode{Label: "a"}))
}

// lookupMaker implements LookupCodeMaker for testing AsLookupCode.
type lookupMaker struct {
	id   string
	name string
}

func (m lookupMaker) LookupCode() LookupCode {
	return LookupCode{Value: m.id, Label: m.name}
}

func TestAsLookupCode(t *testing.T) {

	result := AsLookupCode(lookupMaker{id: "42", name: "Answer"})
	require.Equal(t, "42", result.Value)
	require.Equal(t, "Answer", result.Label)
}

func TestLookupCode_ID(t *testing.T) {
	code := LookupCode{Value: "the-value", Label: "the-label"}
	require.Equal(t, "the-value", code.ID())
}

func TestParseLookupCode_Types(t *testing.T) {

	// LookupCode passes through unchanged
	original := LookupCode{Value: "v", Label: "l"}
	require.Equal(t, original, ParseLookupCode(original))

	// String uses the same value for Value and Label
	require.Equal(t, LookupCode{Value: "x", Label: "x"}, ParseLookupCode("x"))

	// mapof.Any
	fromAny := ParseLookupCode(mapof.Any{"value": "v", "label": "Label", "group": "g"})
	require.Equal(t, "v", fromAny.Value)
	require.Equal(t, "Label", fromAny.Label)
	require.Equal(t, "g", fromAny.Group)

	// mapof.String
	fromString := ParseLookupCode(mapof.String{"value": "v2", "label": "L2"})
	require.Equal(t, "v2", fromString.Value)

	// plain maps
	require.Equal(t, "v3", ParseLookupCode(map[string]any{"value": "v3"}).Value)
	require.Equal(t, "v4", ParseLookupCode(map[string]string{"value": "v4"}).Value)

	// Unknown type returns an empty LookupCode
	require.Equal(t, LookupCode{}, ParseLookupCode(42))
}

func TestLookupCode_GetPointer(t *testing.T) {

	code := &LookupCode{}

	for _, name := range []string{"value", "label", "description", "icon", "group", "href"} {
		pointer, ok := code.GetPointer(name)
		require.True(t, ok, name)
		require.NotNil(t, pointer, name)
	}

	// Unknown property
	_, ok := code.GetPointer("unknown")
	require.False(t, ok)
}

func TestGetLookupCodes_FromSchemaEnum(t *testing.T) {

	// A schema String with an Enum drives getSchemaEnumeration
	element := &Element{Path: "color", Options: mapof.Any{}}
	schemaElement := schema.String{Enum: []string{"red", "green", "blue"}}

	codes, writable := GetLookupCodes(element, schemaElement, nil)
	require.False(t, writable)
	require.Equal(t, 3, len(codes))
	require.Equal(t, "red", codes[0].Value)
}

func TestGetLookupCodes_FromEnumOption(t *testing.T) {

	element := &Element{
		Path:    "color",
		Options: mapof.Any{"enum": "red,green,blue"},
	}

	codes, writable := GetLookupCodes(element, nil, nil)
	require.False(t, writable)
	require.Equal(t, 3, len(codes))
	require.Equal(t, "green", codes[1].Value)
}

func TestParse_Types(t *testing.T) {

	// Element passes through
	original := Element{Type: "test", Path: "name"}
	parsed, err := Parse(original)
	require.NoError(t, err)
	require.Equal(t, original, parsed)

	// map[string]any
	fromMap, err := Parse(map[string]any{"type": "text", "path": "email"})
	require.NoError(t, err)
	require.Equal(t, "text", fromMap.Type)
	require.Equal(t, "email", fromMap.Path)

	// JSON bytes
	fromBytes, err := Parse([]byte(`{"type":"text","path":"phone"}`))
	require.NoError(t, err)
	require.Equal(t, "phone", fromBytes.Path)

	// JSON string
	fromString, err := Parse(`{"type":"text","path":"address"}`)
	require.NoError(t, err)
	require.Equal(t, "address", fromString.Path)

	// Unknown type errors
	_, err = Parse(42)
	require.Error(t, err)

	// Invalid JSON errors
	_, err = Parse([]byte("not json"))
	require.Error(t, err)
}

func TestMustParse(t *testing.T) {

	result := MustParse(map[string]any{"type": "text", "path": "name"})
	require.Equal(t, "text", result.Type)

	// MustParse panics on an unparseable value
	require.Panics(t, func() {
		MustParse(42)
	})
}
