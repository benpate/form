package widget

import (
	"testing"

	"github.com/benpate/form"
	"github.com/stretchr/testify/require"
)

func TestInteger(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type: "text",
			Path: "age",
		},
	)

	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<input name="age" id="age.text" tabIndex="0" type="number" step="1" min="10" max="100" required="true" value="">`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestFloat(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type: "text",
			Path: "distance",
		},
	)

	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<input name="distance" id="distance.text" tabIndex="0" type="number" step="0.01" min="10" max="100" required="true" value="">`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestText(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type: "text",
			Path: "name",
		},
	)

	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<input name="name" id="name.text" tabIndex="0" type="text" maxlength="50" value="">`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestDescription(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(), form.Element{
			Type:        "text",
			Path:        "name",
			Label:       "Widget Label Here... uwu",
			Description: "Hint text no longer added to widgets",
		},
	)

	// schema := getTestSchema()
	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<input name="name" id="name.text" aria-label="Widget Label Here... uwu" aria-description="Hint text no longer added to widgets" tabIndex="0" type="text" maxlength="50" value="">`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestTextTags(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type: "text",
			Path: "tags",
		},
	)

	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<input name="tags" id="tags.text" tabIndex="0" list="datalist-tags" type="text" value=""><datalist id="datalist-tags"><option value="pretty"><option value="please"><option value="my"><option value="dear"><option value="aunt"><option value="sally"></datalist>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestTextTagsWithID(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type: "text",
			Path: "tags",
		},
	)

	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<input name="tags" id="tags.text" tabIndex="0" list="datalist-tags" type="text" value=""><datalist id="datalist-tags"><option value="pretty"><option value="please"><option value="my"><option value="dear"><option value="aunt"><option value="sally"></datalist>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestTextOptions(t *testing.T) {

	UseAll()

	f := form.New(
		getTestSchema(),
		form.Element{
			Type: "text",
			Path: "tag",
			ID:   "tag",
			Options: map[string]any{
				"provider": "test",
			},
		},
	)

	result, err := f.Editor(nil, testLookupProvider{})
	expected := `<input name="tag" id="tag" tabIndex="0" list="datalist-tag" type="text" value=""><datalist id="datalist-tag"><option value="ONE"><option value="TWO"><option value="THREE"><option value="FOUR"><option value="FIVE"></datalist>`

	require.Nil(t, err)
	require.Equal(t, expected, result)
}
