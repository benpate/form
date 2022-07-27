package form

import (
	"testing"

	"github.com/benpate/rosetta/maps"
	"github.com/stretchr/testify/require"
)

func TestToggle(t *testing.T) {

	form := Element{
		Type: "toggle",
		Path: "terms",
	}

	{
		html, err := form.HTML(nil, getTestSchema(), testLookupProvider{})
		require.Nil(t, err)
		require.Equal(t, `<span data-script="install toggle" name="terms"></span>`, html)
		t.Log(html)
	}

	{
		value := maps.Map{"terms": "true"}
		html, err := form.HTML(value, getTestSchema(), testLookupProvider{})
		require.Nil(t, err)
		require.Equal(t, `<span data-script="install toggle" name="terms" value="true"></span>`, html)
		t.Log(html)
	}
}
