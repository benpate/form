package widget

import (
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/stretchr/testify/require"
)

func TestText_SchemaPattern(t *testing.T) {

	UseAll()

	f := form.New(getTestSchema(), form.Element{Type: "text", Path: "code"})

	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<input name="code" id="code.text" tabIndex="0" type="text" maxlength="3" pattern="[A-Z]{3}" value="">`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestText_SchemaPatternOverridesOption(t *testing.T) {

	UseAll()

	f := form.New(getTestSchema(), form.Element{
		Type:    "text",
		Path:    "code",
		Options: mapof.Any{"pattern": "ignored"},
	})

	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<input name="code" id="code.text" tabIndex="0" type="text" maxlength="3" pattern="[A-Z]{3}" value="">`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestText_OptionPatternFallback(t *testing.T) {

	UseAll()

	f := form.New(getTestSchema(), form.Element{
		Type:    "text",
		Path:    "name",
		Options: mapof.Any{"pattern": "[a-z]+"},
	})

	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<input name="name" id="name.text" tabIndex="0" type="text" maxlength="50" pattern="[a-z]+" value="">`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestTextArea_SchemaPattern(t *testing.T) {

	UseAll()

	f := form.New(getTestSchema(), form.Element{Type: "textarea", Path: "code"})

	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<textarea name="code" id="code.textarea" aria-labelledby="code.textarea.label" aria-describedby="code.textarea.description" tabIndex="0" pattern="[A-Z]{3}" maxlength="3"></textarea>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}
