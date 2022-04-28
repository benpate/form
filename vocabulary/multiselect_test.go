package vocabulary

import (
	"testing"

	"github.com/benpate/form"
	"github.com/stretchr/testify/require"
)

func TestMultiselect(t *testing.T) {

	library := getTestLibrary()
	s := getTestSchema()

	f := form.Form{
		Kind: "multiselect",
		Path: "tags",
	}

	value := map[string]interface{}{
		"tags": []string{"pretty", "please"},
	}

	html, err := f.HTML(&library, s, value)

	require.Nil(t, err)
	t.Log(html)
}
