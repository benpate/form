package form

import (
	"testing"

	"github.com/benpate/rosetta/maps"
	"github.com/stretchr/testify/require"
)

func TestSelectOne(t *testing.T) {

	form := Element{
		Type: "select",
		Path: "color",
	}

	html, err := form.HTML(nil, getTestSchema(), testLookupProvider{})

	require.Nil(t, err)
	t.Log(html)
}

func TestSelectOne_WithEnum(t *testing.T) {

	form := Element{
		Type: "select",
		Path: "color",
	}

	value := maps.Map{"color": "Blue"}
	html, err := form.HTML(value, getTestSchema(), testLookupProvider{})

	require.Nil(t, err)
	t.Log(html)
}

func TestSelectOneFromProvider(t *testing.T) {

	form := Element{
		Type:    "select",
		Path:    "color",
		Options: maps.Map{"provider": "test"},
	}

	value := maps.Map{"color": "FOUR"}
	html, err := form.HTML(value, getTestSchema(), testLookupProvider{})

	require.Nil(t, err)
	t.Log(html)
}

func TestSelectOneRadio(t *testing.T) {

	form := Element{
		Type: "radio",
		Path: "color",
	}

	html, err := form.HTML(nil, getTestSchema(), testLookupProvider{})

	require.Nil(t, err)
	t.Log(html)
}

func TestSelectMany(t *testing.T) {

	form := Element{
		Type: "select",
		Path: "tags",
	}

	value := maps.Map{"tags": []string{"pretty", "please"}}

	html, err := form.HTML(value, getTestSchema(), testLookupProvider{})

	require.Nil(t, err)
	t.Log(html)
}
