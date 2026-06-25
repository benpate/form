package form

import (
	"encoding/json"
	"testing"

	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/null"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

// richForm builds a Form exercising every Form field, a wide range of schema
// element types and constraints, and a nested element tree with options.
func richForm() Form {
	return New(
		schema.New(schema.Object{
			Properties: schema.ElementMap{
				"name":      schema.String{MinLength: 1, MaxLength: 100, Required: true},
				"nickname":  schema.String{Format: "no-html", Enum: []string{"a", "b", "c"}},
				"email":     schema.String{Format: "email", RequiredIf: "showEmail is true"},
				"age":       schema.Integer{Minimum: null.NewInt64(0), Maximum: null.NewInt64(150)},
				"score":     schema.Number{Minimum: null.NewFloat(0), Maximum: null.NewFloat(1), BitSize: 64},
				"showEmail": schema.Boolean{Default: null.NewBool(false)},
				"tags":      schema.Array{Items: schema.String{}, MaxLength: 10},
			},
		}),
		Element{
			Type:        "layout-vertical",
			ID:          "root",
			Label:       "Profile",
			Description: "Edit your profile",
			Children: []Element{
				{Type: "text", Path: "name", Label: "Name"},
				{Type: "select", Path: "nickname", Options: mapof.Any{"enum": []any{"a", "b", "c"}}},
				{Type: "text", Path: "email", Options: mapof.Any{"show-if": "showEmail is true"}},
				{Type: "text", Path: "age", Label: "Age"},
				{Type: "text", Path: "score"},
				{Type: "toggle", Path: "showEmail", Label: "Show Email?"},
				{Type: "multiselect", Path: "tags", ReadOnly: true},
			},
		},
		"delivery:fast",
		"layout:vertical",
	)
}

func TestForm_JSONRoundTrip(t *testing.T) {

	original := richForm()

	// Marshal the original form
	first, err := json.Marshal(original)
	require.Nil(t, err)

	// Unmarshal it back into a new Form (this also runs Validate)
	var restored Form
	err = json.Unmarshal(first, &restored)
	require.Nil(t, err)

	// Marshal the restored form and confirm the bytes are identical
	second, err := json.Marshal(restored)
	require.Nil(t, err)

	require.JSONEq(t, string(first), string(second), "marshal -> unmarshal -> marshal must be stable")
}

func TestForm_JSONRoundTrip_FieldByField(t *testing.T) {

	original := richForm()

	data, err := json.Marshal(original)
	require.Nil(t, err)

	var restored Form
	require.Nil(t, json.Unmarshal(data, &restored))

	// Options survive exactly
	require.Equal(t, original.Options, restored.Options)

	// Element tree survives exactly (compare via canonical JSON of just the element)
	requireElementEqual(t, original.Element, restored.Element)

	// Schema survives: compare canonical JSON of each schema
	originalSchema, err := json.Marshal(original.Schema)
	require.Nil(t, err)
	restoredSchema, err := json.Marshal(restored.Schema)
	require.Nil(t, err)
	require.JSONEq(t, string(originalSchema), string(restoredSchema))

	// Spot-check that schema constraints survived the trip
	nameElement, ok := restored.Schema.GetElement("name")
	require.True(t, ok)
	nameString, ok := nameElement.(schema.String)
	require.True(t, ok)
	require.Equal(t, 1, nameString.MinLength)
	require.Equal(t, 100, nameString.MaxLength)
	require.True(t, nameString.Required)

	emailElement, ok := restored.Schema.GetElement("email")
	require.True(t, ok)
	emailString, ok := emailElement.(schema.String)
	require.True(t, ok)
	require.Equal(t, "showEmail is true", emailString.RequiredIf)
}

func TestForm_UnmarshalJSON_ValidationFails(t *testing.T) {

	// The "show-if" references a field ("missing_field") that is not in the schema,
	// so Unmarshal must surface an error.
	data := []byte(`{
		"schema": {
			"type": "object",
			"properties": {
				"name": {"type": "string"}
			}
		},
		"form": {
			"type": "layout-vertical",
			"children": [
				{"type": "text", "path": "name", "options": {"show-if": "missing_field is true"}}
			]
		}
	}`)

	var form Form
	err := json.Unmarshal(data, &form)
	require.Error(t, err)
}

func TestForm_UnmarshalJSON_PathNotInSchema(t *testing.T) {

	// The form references a Path ("not_in_schema") that does not exist in the schema.
	data := []byte(`{
		"schema": {
			"type": "object",
			"properties": {
				"name": {"type": "string"}
			}
		},
		"form": {
			"type": "layout-vertical",
			"children": [
				{"type": "text", "path": "not_in_schema"}
			]
		}
	}`)

	var form Form
	err := json.Unmarshal(data, &form)
	require.Error(t, err)
}

func TestForm_UnmarshalJSON_Valid(t *testing.T) {

	data := []byte(`{
		"schema": {
			"type": "object",
			"properties": {
				"name": {"type": "string"},
				"showEmail": {"type": "boolean"},
				"email": {"type": "string"}
			}
		},
		"form": {
			"type": "layout-vertical",
			"children": [
				{"type": "text", "path": "name"},
				{"type": "toggle", "path": "showEmail"},
				{"type": "text", "path": "email", "options": {"show-if": "showEmail is true"}}
			]
		},
		"options": ["delivery:fast"]
	}`)

	var form Form
	require.Nil(t, json.Unmarshal(data, &form))
	require.Equal(t, []string{"delivery:fast"}, form.Options)
	require.Equal(t, "layout-vertical", form.Element.Type)
	require.Len(t, form.Element.Children, 3)

	// Schema came through
	_, ok := form.Schema.GetElement("showEmail")
	require.True(t, ok)
}

func TestForm_UnmarshalJSON_InvalidJSON(t *testing.T) {

	var form Form
	err := json.Unmarshal([]byte(`{not valid json`), &form)
	require.Error(t, err)
}

// requireElementEqual compares two elements by their canonical JSON form.
func requireElementEqual(t *testing.T, expected Element, actual Element) {
	t.Helper()
	expectedJSON, err := json.Marshal(expected)
	require.Nil(t, err)
	actualJSON, err := json.Marshal(actual)
	require.Nil(t, err)
	require.JSONEq(t, string(expectedJSON), string(actualJSON))
}
