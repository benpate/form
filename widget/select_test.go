package widget

import (
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/sliceof"
	"github.com/stretchr/testify/require"
)

func TestSelectOne(t *testing.T) {

	UseAll()

	element := form.Element{
		Type: "select",
		Path: "color",
	}

	schema := getTestSchema()
	builder := html.New()
	err := element.Edit(&schema, testLookupProvider{}, nil, builder)
	expected := `<select id="select-color" name="color" tabIndex="0"><option></option><option value="Yellow">Yellow</option><option value="Orange">Orange</option><option value="Red">Red</option><option value="Violet">Violet</option><option value="Blue">Blue</option><option value="Green">Green</option></select>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}

func TestSelectOne_ReadOnly(t *testing.T) {

	UseAll()

	element := form.Element{
		Type: "select",
		Path: "color",
		Options: mapof.Any{
			"enum": []form.LookupCode{
				{Value: "YELLOW", Label: "Yellow"},
				{Value: "ORANGE", Label: "Orange"},
				{Value: "RED", Label: "Red"},
				{Value: "VIOLET", Label: "Violet"},
			},
		},
		ReadOnly: true,
	}

	value := mapof.String{
		"color": "RED",
	}

	schema := getTestSchema()
	builder := html.New()
	err := element.View(&schema, nil, value, builder)
	expected := `<div class="layout-value">Red</div>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}

func TestSelectOne_WithEnum(t *testing.T) {

	UseAll()

	schema := schema.Schema{
		Element: schema.Object{
			Properties: schema.ElementMap{
				"data": schema.Object{
					Properties: schema.ElementMap{
						"color": schema.String{Required: false, Enum: []string{"Yellow", "Orange", "Red", "Violet", "Blue", "Green"}},
					},
				},
			},
		},
	}

	element := form.Element{
		Type: "select",
		Path: "data.color",
	}

	value := mapof.Any{"data": mapof.Any{"color": "Blue"}}

	builder := html.New()
	err := element.Edit(&schema, nil, &value, builder)
	expected := `<select id="select-data.color" name="data.color" tabIndex="0"><option></option><option value="Yellow">Yellow</option><option value="Orange">Orange</option><option value="Red">Red</option><option value="Violet">Violet</option><option value="Blue" selected="true">Blue</option><option value="Green">Green</option></select>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}

func TestSelectOneFromProvider(t *testing.T) {

	UseAll()

	element := form.Element{
		Type:    "select",
		Path:    "color",
		Options: map[string]any{"provider": "test"},
	}

	schema := getTestSchema()
	value := mapof.Any{"color": "FOUR"}
	builder := html.New()
	err := element.Edit(&schema, testLookupProvider{}, value, builder)
	expected := `<select id="select-color" name="color" tabIndex="0"><option></option><option value="ONE">This is the first code</option><option value="TWO">This is the second code</option><option value="THREE">This is the third code</option><option value="FOUR" selected="true">This is the fourth code</option><option value="FIVE">This is the fifth code</option></select>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}

func TestSelectOneRadio(t *testing.T) {

	UseAll()

	element := form.Element{
		Type: "radio",
		Path: "color",
	}

	schema := getTestSchema()
	builder := html.New()
	err := element.Edit(&schema, testLookupProvider{}, nil, builder)
	expected := `<label for="radio-color-Yellow"><input type="radio" name="color" id="radio-color-Yellow" value="Yellow">Yellow</label><label for="radio-color-Orange"><input type="radio" name="color" id="radio-color-Orange" value="Orange">Orange</label><label for="radio-color-Red"><input type="radio" name="color" id="radio-color-Red" value="Red">Red</label><label for="radio-color-Violet"><input type="radio" name="color" id="radio-color-Violet" value="Violet">Violet</label><label for="radio-color-Blue"><input type="radio" name="color" id="radio-color-Blue" value="Blue">Blue</label><label for="radio-color-Green"><input type="radio" name="color" id="radio-color-Green" value="Green">Green</label>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}

func TestSelectMany(t *testing.T) {

	UseAll()

	element := form.Element{
		Type: "select",
		Path: "tags",
	}

	value := mapof.Any{"tags": sliceof.String{"pretty", "please"}}

	schema := getTestSchema()

	builder := html.New()
	err := element.Edit(&schema, testLookupProvider{}, &value, builder)
	expected := `<select id="select-tags" name="tags" tabIndex="0"><option></option><option value="pretty" selected="true">pretty</option><option value="please" selected="true">please</option><option value="my">my</option><option value="dear">dear</option><option value="aunt">aunt</option><option value="sally">sally</option></select>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}
