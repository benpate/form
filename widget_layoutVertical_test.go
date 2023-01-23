package form

import (
	"testing"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/mapof"
	"github.com/stretchr/testify/require"
)

func TestLayoutVertical(t *testing.T) {

	element := Element{
		Type:  "layout-vertical",
		Label: "This is my Vertical Layout",
		Children: []Element{
			{
				Type:  "text",
				Label: "Name",
				Path:  "name",
			},
			{
				Type:  "text",
				Label: "Email",
				Path:  "email",
			},
			{
				Type:  "text",
				Label: "Age",
				Path:  "age",
			},
		},
	}

	value := mapof.Any{
		"name":  "John Connor",
		"email": "john@resistance.mil",
		"age":   27,
	}

	schema := getTestSchema()
	builder := html.New()
	err := element.Edit(&schema, nil, &value, builder)
	expected := `<div class="layout layout-vertical"><div class="layout-title">This is my Vertical Layout</div><div class="layout-vertical-elements"><div class="layout-vertical-element"><label>Name</label><input name="name" id="text-name" value="John Connor" type="text" maxlength="50" tabIndex="0"></div><div class="layout-vertical-element"><label>Email</label><input name="email" id="text-email" value="john@resistance.mil" type="email" minlength="10" maxlength="100" required="true" tabIndex="0"></div><div class="layout-vertical-element"><label>Age</label><input name="age" id="text-age" value="27" type="number" step="1" min="10" max="100" required="true" tabIndex="0"></div></div></div>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}

func TestRules(t *testing.T) {

	form := Element{
		Type:  "layout-vertical",
		Label: "This is my Vertical Layout",
		Children: []Element{
			{
				Type:  "text",
				Label: "Name",
				Path:  "name",
			},
			{
				Type:  "text",
				Label: "Email",
				Path:  "email",
			},
			{
				Type:  "text",
				Label: "Age",
				Path:  "age",
			},
		},
	}

	builder := html.New()
	schema := getTestSchema()
	err := form.Edit(&schema, nil, nil, builder)
	expected := `<div class="layout layout-vertical"><div class="layout-title">This is my Vertical Layout</div><div class="layout-vertical-elements"><div class="layout-vertical-element"><label>Name</label><input name="name" id="text-name" type="text" maxlength="50" tabIndex="0"></div><div class="layout-vertical-element"><label>Email</label><input name="email" id="text-email" type="email" minlength="10" maxlength="100" required="true" tabIndex="0"></div><div class="layout-vertical-element"><label>Age</label><input name="age" id="text-age" type="number" step="1" min="10" max="100" required="true" tabIndex="0"></div></div></div>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}
