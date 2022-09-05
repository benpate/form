package form

import (
	"testing"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/maps"
	"github.com/stretchr/testify/require"
)

func TestToggle(t *testing.T) {

	form := Element{
		Type: "toggle",
		Path: "terms",
	}

	{
		builder := html.New()
		schema := getTestSchema()
		err := form.Edit(&schema, testLookupProvider{}, nil, builder)
		expected := `<span data-script="install toggle" name="terms"></span>`

		require.Nil(t, err)
		require.Equal(t, expected, builder.String())
	}

	{
		value := maps.Map{"terms": "true"}
		builder := html.New()
		schema := getTestSchema()
		err := form.Edit(&schema, testLookupProvider{}, value, builder)
		expected := `<span data-script="install toggle" name="terms" value="true"></span>`

		require.Nil(t, err)
		require.Equal(t, expected, builder.String())
	}
}
