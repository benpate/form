package widget

import (
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/stretchr/testify/require"
)

func TestTextarea(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type:    "textarea",
			Path:    "username",
			Options: mapof.Any{"pattern": "[a-z]+"},
		},
	)

	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<textarea name="username" id="username.textarea" aria-labelledby="username.textarea.label" aria-describedby="username.textarea.description" tabIndex="0" pattern="[a-z]+" minlength="10" maxlength="100" required="true"></textarea>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestTextareaRows(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type:    "textarea",
			Path:    "username",
			Options: mapof.Any{"rows": 4.0},
		},
	)

	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<textarea name="username" id="username.textarea" rows="4" aria-labelledby="username.textarea.label" aria-describedby="username.textarea.description" tabIndex="0" minlength="10" maxlength="100" required="true"></textarea>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}
