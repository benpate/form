package widget

import (
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/stretchr/testify/require"
)

func TestTextarea(t *testing.T) {

	UseAll()

	element := form.Element{
		Type: "textarea",
		Path: "username",
	}

	builder := html.New()
	schema := getTestSchema()
	err := element.Edit(&schema, testLookupProvider{}, nil, builder)
	expected := `<textarea name="username" id="username.textarea" aria-labelledby="username.textarea.label" aria-describedby="username.textarea.description" tabIndex="0" minlength="10" maxlength="100" pattern="[a-z]+" required="true"></textarea>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}
