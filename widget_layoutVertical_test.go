package form

import (
	"testing"

	"github.com/benpate/derp"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/maps"

	"github.com/stretchr/testify/assert"
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

	value := maps.Map{
		"name":  "John Connor",
		"email": "john@resistance.mil",
		"age":   27,
	}

	schema := getTestSchema()
	builder := html.New()
	err := element.Edit(&schema, nil, value, builder)

	assert.Nil(t, err)
	t.Log(builder.String())
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

	assert.Nil(t, err)
	derp.Report(err)
	t.Log(builder.String())
}
