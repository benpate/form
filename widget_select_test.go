package form

import (
	"testing"

	"github.com/benpate/rosetta/maps"
	"github.com/stretchr/testify/require"
)

func TestSelectOne(t *testing.T) {

	form := Element{
		Type: "select",
		Path: "color",
	}

	actual, err := form.HTML(nil, getTestSchema(), testLookupProvider{})
	expected := `<select id="select-color" name="color" tabIndex="0"><option></option><option value="Yellow">Yellow</option><option value="Orange">Orange</option><option value="Red">Red</option><option value="Violet">Violet</option><option value="Blue">Blue</option><option value="Green">Green</option></select>`

	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

func TestSelectOne_WithEnum(t *testing.T) {

	form := Element{
		Type: "select",
		Path: "color",
	}

	value := maps.Map{"color": "Blue"}
	expected, err := form.HTML(value, getTestSchema(), testLookupProvider{})
	actual := `<select id="select-color" name="color" tabIndex="0"><option></option><option value="Yellow">Yellow</option><option value="Orange">Orange</option><option value="Red">Red</option><option value="Violet">Violet</option><option value="Blue" selected="true">Blue</option><option value="Green">Green</option></select>`

	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

func TestSelectOneFromProvider(t *testing.T) {

	form := Element{
		Type:    "select",
		Path:    "color",
		Options: maps.Map{"provider": "test"},
	}

	value := maps.Map{"color": "FOUR"}
	actual, err := form.HTML(value, getTestSchema(), testLookupProvider{})
	expected := `<select id="select-color" name="color" tabIndex="0"><option></option><option value="ONE">This is the first code</option><option value="TWO">This is the second code</option><option value="THREE">This is the third code</option><option value="FOUR" selected="true">This is the fourth code</option><option value="FIVE">This is the fifth code</option></select>`

	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

func TestSelectOneRadio(t *testing.T) {

	form := Element{
		Type: "radio",
		Path: "color",
	}

	actual, err := form.HTML(nil, getTestSchema(), testLookupProvider{})
	expected := `<label for="radio-color-Yellow"><input type="radio" name="color" id="radio-color-Yellow" value="Yellow">Yellow</label><label for="radio-color-Orange"><input type="radio" name="color" id="radio-color-Orange" value="Orange">Orange</label><label for="radio-color-Red"><input type="radio" name="color" id="radio-color-Red" value="Red">Red</label><label for="radio-color-Violet"><input type="radio" name="color" id="radio-color-Violet" value="Violet">Violet</label><label for="radio-color-Blue"><input type="radio" name="color" id="radio-color-Blue" value="Blue">Blue</label><label for="radio-color-Green"><input type="radio" name="color" id="radio-color-Green" value="Green">Green</label>`

	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

func TestSelectMany(t *testing.T) {

	form := Element{
		Type: "select",
		Path: "tags",
	}

	value := maps.Map{"tags": []string{"pretty", "please"}}

	actual, err := form.HTML(value, getTestSchema(), testLookupProvider{})
	expected := `<select id="select-tags" name="tags" tabIndex="0"><option></option><option value="pretty" selected="true">pretty</option><option value="please">please</option><option value="my">my</option><option value="dear">dear</option><option value="aunt">aunt</option><option value="sally">sally</option></select>`

	require.Nil(t, err)
	require.Equal(t, expected, actual)
}
