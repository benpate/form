package widget

import (
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/stretchr/testify/require"
)

func TestInteger(t *testing.T) {

	UseAll()

	element := form.Element{
		Type: "text",
		Path: "age",
	}

	builder := html.New()
	schema := getTestSchema()
	err := element.Edit(&schema, testLookupProvider{}, nil, builder)
	expected := `<input name="age" id="text-age" type="number" step="1" min="10" max="100" required="true" tabIndex="0">`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}

func TestFloat(t *testing.T) {

	UseAll()

	element := form.Element{
		Type: "text",
		Path: "distance",
	}

	builder := html.New()
	schema := getTestSchema()
	err := element.Edit(&schema, testLookupProvider{}, nil, builder)
	expected := `<input name="distance" id="text-distance" type="number" step="0.01" min="10" max="100" required="true" tabIndex="0">`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}

func TestText(t *testing.T) {

	UseAll()

	element := form.Element{
		Type: "text",
		Path: "username",
	}

	builder := html.New()
	schema := getTestSchema()
	err := element.Edit(&schema, testLookupProvider{}, nil, builder)
	expected := `<input name="username" id="text-username" type="text" minlength="10" maxlength="100" pattern="[a-z]+" required="true" tabIndex="0">`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}

func TestDescription(t *testing.T) {

	UseAll()

	element := form.Element{
		Type:        "text",
		Path:        "name",
		Description: "Hint text no longer added to widgets",
	}

	builder := html.New()
	schema := getTestSchema()
	err := element.Edit(&schema, testLookupProvider{}, nil, builder)
	expected := `<input name="name" id="text-name" type="text" maxlength="50" tabIndex="0">`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}

func TestTextTags(t *testing.T) {

	UseAll()

	element := form.Element{
		Type: "text",
		Path: "tags",
	}

	builder := html.New()
	schema := getTestSchema()
	err := element.Edit(&schema, testLookupProvider{}, nil, builder)
	expected := `<input name="tags" id="text-tags" list="datalist-tags" type="text" tabIndex="0"><datalist id="datalist-tags"><option value="pretty"><option value="please"><option value="my"><option value="dear"><option value="aunt"><option value="sally"></datalist>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}

func TestTextTagsWithID(t *testing.T) {

	UseAll()

	element := form.Element{
		Type: "text",
		Path: "tags",
	}

	builder := html.New()
	schema := getTestSchema()
	err := element.Edit(&schema, testLookupProvider{}, nil, builder)
	expected := `<input name="tags" id="text-tags" list="datalist-tags" type="text" tabIndex="0"><datalist id="datalist-tags"><option value="pretty"><option value="please"><option value="my"><option value="dear"><option value="aunt"><option value="sally"></datalist>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}

func TestTextOptions(t *testing.T) {

	UseAll()

	element := form.Element{
		Type: "text",
		Path: "tag",
		ID:   "tag",
		Options: map[string]any{
			"provider": "test",
		},
	}

	builder := html.New()
	schema := getTestSchema()
	err := element.Edit(&schema, testLookupProvider{}, nil, builder)
	expected := `<input name="tag" id="tag" list="datalist-tag" type="text" tabIndex="0"><datalist id="datalist-tag"><option value="ONE"><option value="TWO"><option value="THREE"><option value="FOUR"><option value="FIVE"></datalist>`

	require.Nil(t, err)
	require.Equal(t, expected, builder.String())
}
