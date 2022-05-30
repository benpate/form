package vocabulary

import (
	"testing"

	"github.com/benpate/form"
	"github.com/stretchr/testify/require"
)

func TestSelectOne(t *testing.T) {

	library := getTestLibrary()
	s := getTestSchema()

	f := form.Form{
		Kind: "select",
		Path: "color",
	}

	html, err := f.HTML(&library, s, nil)

	require.Nil(t, err)
	t.Log(html)
}

func TestSelectOne_WithEnum(t *testing.T) {

	library := getTestLibrary()
	s := getTestSchema()

	f := form.Form{
		Kind: "select",
		Path: "color",
	}

	value := map[string]any{"color": "Blue"}

	html, err := f.HTML(&library, s, value)

	require.Nil(t, err)
	t.Log(html)
}

func TestSelectOneFromProvider(t *testing.T) {

	library := getTestLibrary()
	s := getTestSchema()

	f := form.Form{
		Kind:    "select",
		Path:    "color",
		Options: map[string]any{"provider": "/test"},
	}

	value := map[string]any{"color": "FOUR"}

	html, err := f.HTML(&library, s, value)

	require.Nil(t, err)
	t.Log(html)
}

func TestSelectOneRadio(t *testing.T) {

	library := getTestLibrary()
	s := getTestSchema()

	f := form.Form{
		Kind: "select",
		Path: "color",
		Options: map[string]any{
			"format": "radio",
		},
	}

	html, err := f.HTML(&library, s, nil)

	require.Nil(t, err)
	t.Log(html)
}

func TestSelectMany(t *testing.T) {

	library := getTestLibrary()
	s := getTestSchema()

	f := form.Form{
		Kind: "select",
		Path: "tags",
	}

	value := map[string]any{
		"tags": []string{"pretty", "please"},
	}

	html, err := f.HTML(&library, s, value)

	require.Nil(t, err)
	t.Log(html)
}
