package form

import (
	"encoding/json"
	"testing"

	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

// TestElement_UnmarshalMap_ReadOnly pins the behavior of convert.Bool on the
// "readOnly" field, which is read from an untyped map[string]any (e.g. a JSON or
// HJSON form definition). rosetta v0.27.0 changed convert.Bool's numeric path so
// that only 0 and 1 are recognized (1 -> true, every other number -> false),
// where it previously treated any non-zero number as true. This test locks the
// current contract for every realistic readOnly representation.
func TestElement_UnmarshalMap_ReadOnly(t *testing.T) {

	assert := func(t *testing.T, readOnlyValue any, expected bool) {
		t.Helper()

		element := NewElement()
		err := element.UnmarshalMap(map[string]any{
			"type":     "text",
			"path":     "name",
			"readOnly": readOnlyValue,
		})

		require.NoError(t, err)
		require.Equal(t, expected, element.ReadOnly,
			"readOnly input %#v (%T) should unmarshal to ReadOnly=%v", readOnlyValue, readOnlyValue, expected)
	}

	// Native bool values: the canonical, always-correct representation.
	assert(t, true, true)
	assert(t, false, false)

	// String values, as they appear in HJSON form definitions.
	assert(t, "true", true)
	assert(t, "false", false)
	assert(t, "", false)

	// Numeric 0/1 round-trip losslessly to bool and are unchanged across versions.
	assert(t, 1, true)
	assert(t, 0, false)

	// Absent key: missing readOnly means "not read-only".
	assert(t, nil, false)
}

// TestElement_UnmarshalMap_ReadOnly_NumericRegression documents the one v0.27.0
// behavior change that reaches form: a "readOnly" set to a number other than 0 or
// 1. Under rosetta v0.25.36, convert.Bool returned TRUE for any non-zero number;
// under v0.27.0 it returns FALSE for anything that is not exactly 1. JSON decodes
// numbers as float64, which is exactly how such a value would arrive here.
func TestElement_UnmarshalMap_ReadOnly_NumericRegression(t *testing.T) {

	assert := func(t *testing.T, jsonValue string, expected bool) {
		t.Helper()

		var data map[string]any
		require.NoError(t, json.Unmarshal([]byte(`{"type":"text","path":"name","readOnly":`+jsonValue+`}`), &data))

		element := NewElement()
		require.NoError(t, element.UnmarshalMap(data))
		require.Equal(t, expected, element.ReadOnly,
			"readOnly JSON %s should unmarshal to ReadOnly=%v under rosetta v0.27.0", jsonValue, expected)
	}

	// The safe, realistic values behave intuitively.
	assert(t, "true", true)
	assert(t, "false", false)
	assert(t, "1", true)
	assert(t, "0", false)

	// The changed cases: any number other than 0 or 1 is now FALSE (was TRUE in v0.25.36).
	assert(t, "2", false)
	assert(t, "-1", false)
	assert(t, "1.5", false)
}

// nestedGetSchema returns a schema with one level of nesting, used to exercise
// rosetta v0.27.0's rewritten recursive schema.Get path through form's
// Element.getValue / Schema.Get usage.
func nestedGetSchema() schema.Schema {
	return schema.New(schema.Object{
		Properties: schema.ElementMap{
			"name": schema.String{},
			"address": schema.Object{
				Properties: schema.ElementMap{
					"city": schema.String{},
					"zip":  schema.String{},
				},
			},
		},
	})
}

// TestElement_GetValue_NestedPath verifies that form's getValue (which delegates
// to schema.Schema.Get) resolves both top-level and dotted nested paths against a
// mapof.Any object. rosetta v0.27.0 replaced the list-based recursive getter with
// a strings.Cut-based one; this confirms the traversal still reaches nested values
// and returns nil (no panic) for paths that do not exist.
func TestElement_GetValue_NestedPath(t *testing.T) {

	s := nestedGetSchema()

	object := mapof.Any{
		"name": "Sarah",
		"address": mapof.Any{
			"city": "Cupertino",
			"zip":  "95014",
		},
	}

	assert := func(t *testing.T, path string, expected any) {
		t.Helper()
		element := &Element{Type: "text", Path: path}
		require.Equal(t, expected, element.getValue(&object, &s),
			"getValue(%q) should resolve to %#v", path, expected)
	}

	// Top-level path.
	assert(t, "name", "Sarah")

	// Nested paths through the rewritten recursive getter.
	assert(t, "address.city", "Cupertino")
	assert(t, "address.zip", "95014")

	// A path that is not in the schema resolves to nil without panicking.
	assert(t, "address.country", nil)
	assert(t, "missing", nil)
}

// TestElement_GetSchema_NestedPath verifies that Element.GetSchema (which delegates
// to schema.Schema.GetElement) resolves nested element definitions. GetElement's
// signature is unchanged across the upgrade; this pins that nested lookup still works.
func TestElement_GetSchema_NestedPath(t *testing.T) {

	s := nestedGetSchema()

	// A nested path returns its concrete element definition (a String element).
	element := &Element{Type: "text", Path: "address.city"}
	resolved := element.GetSchema(&s)
	require.NotNil(t, resolved)
	_, ok := resolved.(schema.String)
	require.True(t, ok, "address.city should resolve to a schema.String element")

	// A path that is not in the schema resolves to a nil element.
	missing := &Element{Type: "text", Path: "address.country"}
	require.Nil(t, missing.GetSchema(&s))
}
