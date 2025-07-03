package widget

import (
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/stretchr/testify/require"
)

func TestToggle_View(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type: "toggle",
			Path: "terms",
		},
	)

	value := mapof.Any{
		"terms": "true",
	}

	result, err := f.Editor(&value, testLookupProvider{})
	expected := `<span id="toggle-terms-true" data-script="install toggle " name="terms" value="true"></span>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestToggle_Edit(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type: "toggle",
			Path: "terms",
		},
	)

	value := mapof.Any{
		"terms": "true",
	}

	result, err := f.Editor(&value, testLookupProvider{})
	expected := `<span id="toggle-terms-true" data-script="install toggle " name="terms" value="true"></span>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}
