package form

import (
	"testing"

	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

func testValidateSchema() schema.Schema {
	return schema.New(schema.Object{
		Properties: schema.ElementMap{
			"name":      schema.String{},
			"email":     schema.String{},
			"showEmail": schema.Boolean{},
		},
	})
}

func TestFormValidate_Valid(t *testing.T) {

	form := New(
		testValidateSchema(),
		Element{
			Type: "layout-vertical",
			Children: []Element{
				{Type: "text", Path: "name"},
				{Type: "text", Path: "email", Options: mapof.Any{"show-if": "showEmail is true"}},
				{Type: "toggle", Path: "showEmail"},
			},
		},
	)

	require.Nil(t, form.Validate())
}

func TestFormValidate_MissingPath(t *testing.T) {

	form := New(
		testValidateSchema(),
		Element{
			Type: "layout-vertical",
			Children: []Element{
				{Type: "text", Path: "name"},
				{Type: "text", Path: "not_in_schema"},
			},
		},
	)

	require.Error(t, form.Validate())
}

func TestFormValidate_MissingShowIfField(t *testing.T) {

	form := New(
		testValidateSchema(),
		Element{
			Type: "layout-vertical",
			Children: []Element{
				{Type: "text", Path: "email", Options: mapof.Any{"show-if": "missing_field is true"}},
			},
		},
	)

	require.Error(t, form.Validate())
}

func TestFormValidate_NestedChildError(t *testing.T) {

	form := New(
		testValidateSchema(),
		Element{
			Type: "layout-vertical",
			Children: []Element{
				{
					Type: "layout-vertical",
					Children: []Element{
						{Type: "text", Path: "name"},
						{Type: "text", Path: "deeply_missing"},
					},
				},
			},
		},
	)

	require.Error(t, form.Validate())
}

func TestFormValidate_LayoutElementNoPath(t *testing.T) {

	// Layout/container elements legitimately have no Path; they must not trigger an error.
	form := New(
		testValidateSchema(),
		Element{
			Type: "layout-vertical",
			Children: []Element{
				{Type: "text", Path: "name"},
			},
		},
	)

	require.Nil(t, form.Validate())
}
