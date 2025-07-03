package widget

import (
	"testing"

	"github.com/benpate/form"
	"github.com/stretchr/testify/require"
)

func TestTextarea(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type: "textarea",
			Path: "username",
		},
	)

	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<textarea name="username" id="username.textarea" aria-labelledby="username.textarea.label" aria-describedby="username.textarea.description" tabIndex="0" minlength="10" maxlength="100" pattern="[a-z]+" required="true"></textarea>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}
