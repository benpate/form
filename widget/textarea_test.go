package widget

import (
	"testing"

	"github.com/benpate/html"
	"github.com/stretchr/testify/require"
)

func TestTextarea(t *testing.T) {

	element := Element{
		Type: "textarea",
		Path: "username",
	}

	builder := html.New()
	schema := getTestSchema()
	err := element.Edit(&schema, testLookupProvider{}, nil, builder)
	expected := `<textarea name="username" id="textarea-username" minlength="10" maxlength="100" pattern="[a-z]+" required="true" tabIndex="0"></textarea>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}
