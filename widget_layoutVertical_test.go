package form

import (
	"testing"

	"github.com/benpate/derp"
	"github.com/benpate/rosetta/maps"

	"github.com/stretchr/testify/assert"
)

func TestLayoutVertical(t *testing.T) {

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

	value := maps.Map{
		"name":  "John Connor",
		"email": "john@resistance.mil",
		"age":   27,
	}

	html, err := form.HTML(value, getTestSchema(), nil)

	assert.Nil(t, err)
	t.Log(html)
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

	html, err := form.HTML(nil, getTestSchema(), nil)

	assert.Nil(t, err)
	derp.Report(err)
	t.Log(html)
}
