package widget

import (
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/sliceof"
	"github.com/stretchr/testify/require"
)

func TestSelectOne(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type: "select",
			Path: "color",
		},
	)

	// schema := getTestSchema()
	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<select id="select-color" name="color" aria-labelledby=".label" aria-describedby=".description" tabIndex="0"><option></option><option value="Yellow">Yellow</option><option value="Orange">Orange</option><option value="Red">Red</option><option value="Violet">Violet</option><option value="Blue">Blue</option><option value="Green">Green</option></select>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestSelectOne_ReadOnly(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
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
		},
	)

	value := mapof.String{
		"color": "RED",
	}

	result, err := f.Editor(&value, nil)
	expected := `<div class="layout-value">Red</div>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestSelectOne_View(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
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
		},
	)

	value := mapof.String{
		"color": "RED",
	}

	// schema := getTestSchema()
	result, err := f.Viewer(value, nil)
	expected := `<div class="layout-value">Red</div>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestSelectOne_WithGroups(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type: "select",
			Path: "other",
			Options: mapof.Any{
				"enum": []form.LookupCode{
					{Group: "Colour", Value: "YELLOW", Label: "Yellow"},
					{Group: "Colour", Value: "ORANGE", Label: "Orange"},
					{Group: "Colour", Value: "RED", Label: "Red"},
					{Group: "Colour", Value: "VIOLET", Label: "Violet"},
					{Group: "Flavour", Value: "BITTER", Label: "Bitter"},
					{Group: "Flavour", Value: "SALTY", Label: "Salty"},
					{Group: "Flavour", Value: "SOUR", Label: "Sour"},
					{Group: "Flavour", Value: "SWEET", Label: "Sweet"},
					{Group: "Flavour", Value: "UMAMI", Label: "Umami"},
				},
			},
		},
	)

	value := mapof.String{
		"other": "RED",
	}

	result, err := f.Editor(value, nil)
	expected := `<select id="select-other" name="other" aria-labelledby=".label" aria-describedby=".description" tabIndex="0"><option></option><optgroup label="Colour"><option value="YELLOW">Yellow</option><option value="ORANGE">Orange</option><option value="RED" selected="true">Red</option><option value="VIOLET">Violet</option></optgroup><optgroup label="Flavour"><option value="BITTER">Bitter</option><option value="SALTY">Salty</option><option value="SOUR">Sour</option><option value="SWEET">Sweet</option><option value="UMAMI">Umami</option></optgroup></select>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestSelectOne_WithEnum(t *testing.T) {

	UseAll()

	f := form.New(schema.Schema{
		Element: schema.Object{
			Properties: schema.ElementMap{
				"data": schema.Object{
					Properties: schema.ElementMap{
						"color": schema.String{Required: false, Enum: []string{"Yellow", "Orange", "Red", "Violet", "Blue", "Green"}},
					},
				},
			},
		},
	},
		form.Element{
			Type: "select",
			Path: "data.color",
		},
	)

	value := mapof.Any{"data": mapof.Any{"color": "Blue"}}

	result, err := f.Editor(&value, nil)
	expected := `<select id="select-data.color" name="data.color" aria-labelledby=".label" aria-describedby=".description" tabIndex="0"><option></option><option value="Yellow">Yellow</option><option value="Orange">Orange</option><option value="Red">Red</option><option value="Violet">Violet</option><option value="Blue" selected="true">Blue</option><option value="Green">Green</option></select>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestSelectOneFromProvider(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type:    "select",
			Path:    "color",
			Options: map[string]any{"provider": "test"},
		},
	)

	// schema := getTestSchema()
	value := mapof.Any{"color": "FOUR"}
	result, err := f.Editor(&value, testLookupProvider{})
	expected := `<select id="select-color" name="color" aria-labelledby=".label" aria-describedby=".description" tabIndex="0"><option></option><option value="ONE">This is the first code</option><option value="TWO">This is the second code</option><option value="THREE">This is the third code</option><option value="FOUR" selected="true">This is the fourth code</option><option value="FIVE">This is the fifth code</option></select>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestSelectOneRadio(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type: "radio",
			Path: "color",
		},
	)

	// schema := getTestSchema()
	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<label for="color.radio.Yellow" id="color.radio.label"><input type="radio" name="color" id="color.radio.Yellow" value="Yellow" aria-label="Yellow" tabIndex="0"><span aria-hidden="true">Yellow</span></label><label for="color.radio.Orange" id="color.radio.label"><input type="radio" name="color" id="color.radio.Orange" value="Orange" aria-label="Orange" tabIndex="0"><span aria-hidden="true">Orange</span></label><label for="color.radio.Red" id="color.radio.label"><input type="radio" name="color" id="color.radio.Red" value="Red" aria-label="Red" tabIndex="0"><span aria-hidden="true">Red</span></label><label for="color.radio.Violet" id="color.radio.label"><input type="radio" name="color" id="color.radio.Violet" value="Violet" aria-label="Violet" tabIndex="0"><span aria-hidden="true">Violet</span></label><label for="color.radio.Blue" id="color.radio.label"><input type="radio" name="color" id="color.radio.Blue" value="Blue" aria-label="Blue" tabIndex="0"><span aria-hidden="true">Blue</span></label><label for="color.radio.Green" id="color.radio.label"><input type="radio" name="color" id="color.radio.Green" value="Green" aria-label="Green" tabIndex="0"><span aria-hidden="true">Green</span></label>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestSelectMany(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type: "select",
			Path: "tags",
		},
	)

	value := mapof.Any{"tags": sliceof.String{"pretty", "please"}}

	result, err := f.Editor(&value, testLookupProvider{})
	expected := `<select id="select-tags" name="tags" aria-labelledby=".label" aria-describedby=".description" tabIndex="0"><option></option><option value="pretty" selected="true">pretty</option><option value="please" selected="true">please</option><option value="my">my</option><option value="dear">dear</option><option value="aunt">aunt</option><option value="sally">sally</option></select>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}
