package form

import (
	"testing"

	"github.com/benpate/rosetta/maps"
	"github.com/stretchr/testify/require"
)

func TestInteger(t *testing.T) {

	form := Element{
		Type: "text",
		Path: "age",
	}

	html, err := form.HTML(nil, getTestSchema(), testLookupProvider{})

	require.Nil(t, err)
	require.Equal(t, `<input name="age" id="text-age" type="number" step="1" min="10" max="100" required="true" tabIndex="0">`, html)
}

func TestFloat(t *testing.T) {

	form := Element{
		Type: "text",
		Path: "distance",
	}

	html, err := form.HTML(nil, getTestSchema(), testLookupProvider{})

	require.Nil(t, err)
	require.Equal(t, `<input name="distance" id="text-distance" type="number" step="0.01" min="10" max="100" required="true" tabIndex="0">`, html)
}

func TestText(t *testing.T) {

	form := Element{
		Type: "text",
		Path: "username",
	}

	html, err := form.HTML(nil, getTestSchema(), testLookupProvider{})

	require.Nil(t, err)
	require.Equal(t, `<input name="username" id="text-username" type="text" minlength="10" maxlength="100" pattern="[a-z]+" required="true" tabIndex="0">`, html)
}

func TestDescription(t *testing.T) {

	form := Element{
		Type:        "text",
		Path:        "name",
		Description: "Hint text no longer added to widgets",
	}

	html, err := form.HTML(nil, getTestSchema(), testLookupProvider{})

	require.Nil(t, err)
	require.Equal(t, `<input name="name" id="text-name" type="text" maxlength="50" tabIndex="0">`, html)
}

func TestTextTags(t *testing.T) {

	form := Element{
		Type: "text",
		Path: "tags",
	}

	html, err := form.HTML(nil, getTestSchema(), testLookupProvider{})

	require.Nil(t, err)
	require.Equal(t, `<input name="tags" id="text-tags" list="datalist-tags" type="text" tabIndex="0"><datalist id="datalist-tags"><option value="pretty"><option value="please"><option value="my"><option value="dear"><option value="aunt"><option value="sally"></datalist>`, html)
}

func TestTextTagsWithID(t *testing.T) {

	form := Element{
		Type: "text",
		Path: "tags",
	}

	html, err := form.HTML(nil, getTestSchema(), testLookupProvider{})

	require.Nil(t, err)
	require.Equal(t, `<input name="tags" id="text-tags" list="datalist-tags" type="text" tabIndex="0"><datalist id="datalist-tags"><option value="pretty"><option value="please"><option value="my"><option value="dear"><option value="aunt"><option value="sally"></datalist>`, html)
}

func TestTextOptions(t *testing.T) {

	form := Element{
		Type: "text",
		Path: "tag",
		ID:   "tag",
		Options: maps.Map{
			"datasource": "test",
		},
	}

	html, err := form.HTML(nil, getTestSchema(), testLookupProvider{})

	require.Nil(t, err)
	require.Equal(t, `<input name="tag" id="text-tag" list="datalist-tag" type="text" tabIndex="0"><datalist id="datalist-tag"><option value="ONE"><option value="TWO"><option value="THREE"><option value="FOUR"><option value="FIVE"></datalist>`, html)
}
