package vocabulary

import (
	"testing"

	"github.com/benpate/form"
	"github.com/stretchr/testify/require"
)

func TestToggle(t *testing.T) {

	library := getTestLibrary()
	s := getTestSchema()

	f := form.Form{
		Kind: "toggle",
		Path: "terms",
	}

	{
		html, err := f.HTML(&library, s, nil)
		require.Nil(t, err)
		require.Equal(t, `<span data-script="install toggle" name="terms"></span>`, html)
		//t.Log(html)
	}

	{
		html, err := f.HTML(&library, s, map[string]any{"terms": "true"})
		require.Nil(t, err)
		require.Equal(t, `<span data-script="install toggle" name="terms" value="true"></span>`, html)
		// t.Log(html)
	}
}
