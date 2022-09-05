package form

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormEditor(t *testing.T) {

	data := map[string]any{
		"name":  "John Connor",
		"email": "john@connor.mil",
		"age":   42,
		"human": true,
		"ology": map[string]any{
			"biology":    "I am a biological human",
			"psychology": "I think, therefore I am",
			"ontology":   "Here's why I am a human",
		},
	}

	form := getTestForm()

	result, err := form.Viewer(data, nil)
	require.Nil(t, err)
	t.Log(result)
}

func TestFormViewer(t *testing.T) {

	data := map[string]any{
		"name":  "John Connor",
		"email": "john@connor.mil",
		"age":   42,
		"human": true,
		"ology": map[string]any{
			"biology":    "I am a biological human",
			"psychology": "I think, therefore I am",
			"ontology":   "Here's why I am a human",
		},
	}

	form := getTestForm()
	result, err := form.Viewer(data, nil)

	require.Nil(t, err)
	t.Log(result)
}

func TestAllPaths(t *testing.T) {

	form := getTestElement()

	paths := form.AllElements()

	require.Equal(t, 5, len(paths))
	require.Equal(t, "data.firstName", paths[0].Path)
	require.Equal(t, "data.lastName", paths[1].Path)
	require.Equal(t, "data.biography", paths[2].Path)
	require.Equal(t, "data.psychology", paths[3].Path)
	require.Equal(t, "data.ontology", paths[4].Path)
}

func getTestForm() Form {
	return New(getTestSchema(), getTestElement())
}

func getTestElement() Element {

	return Element{
		Type: "layout-vertical",
		Children: []Element{
			{
				Type:  "text",
				ID:    "idName",
				Label: "Name",
				Path:  "name",
			}, {
				Type:  "text",
				ID:    "idEmail",
				Label: "Email",
				Path:  "email",
			}, {
				Type:  "text",
				ID:    "idAge",
				Label: "Age",
				Path:  "age",
			}, {
				Type:  "toggle",
				ID:    "idHuman",
				Label: "Human",
				Path:  "human",
			}, {
				Type: "layout-horizontal",
				Children: []Element{
					{
						Type:  "textarea",
						ID:    "idBiology",
						Label: "Biology",
						Path:  "ology.biology",
					}, {
						Type:  "textarea",
						ID:    "idPsychology",
						Label: "Psychology",
						Path:  "ology.psychology",
					}, {
						Type:  "textarea",
						ID:    "idOntology",
						Label: "Ontology",
						Path:  "ology.ontology",
					},
				},
			},
		},
	}
}
