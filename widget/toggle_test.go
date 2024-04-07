package form

import (
	"testing"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/mapof"
	"github.com/stretchr/testify/require"
)

func TestToggle_View(t *testing.T) {

	form := Element{
		Type: "toggle",
		Path: "terms",
	}

	value := mapof.Any{
		"terms": "true",
	}

	builder := html.New()
	schema := getTestSchema()
	err := form.Edit(&schema, testLookupProvider{}, &value, builder)
	expected := `<span data-script="install toggle" name="terms" value="true"></span>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}

func TestToggle_Edit(t *testing.T) {

	form := Element{
		Type: "toggle",
		Path: "terms",
	}

	value := mapof.Any{
		"terms": "true",
	}

	builder := html.New()
	schema := getTestSchema()
	err := form.Edit(&schema, testLookupProvider{}, &value, builder)
	expected := `<span data-script="install toggle" name="terms" value="true"></span>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}
