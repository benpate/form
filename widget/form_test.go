package widget

import (
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/stretchr/testify/require"
)

func TestFormEditor(t *testing.T) {

	data := mapof.Any{
		"name":  "John Connor",
		"email": "john@connor.mil",
		"age":   42,
		"human": true,
		"ology": mapof.Any{
			"biology":    "I am a biological human",
			"psychology": "I think, therefore I am",
			"ontology":   "Here's why I am a human",
		},
	}

	f := getTestForm()

	_, err := f.Viewer(data, nil)
	// expected := `<div class="layout layout-vertical"><div class="layout-elements"><div class="layout-element"><label for="idName">Name</label><div class="layout-value"></div></div><div class="layout-element"><label for="idEmail">Email</label><div class="layout-value"></div></div><div class="layout-element"><label for="idAge">Age</label><div class="layout-value"></div></div><div class="layout-element"><label for="idHuman">Human</label><div class="layout-value"></div></div><div class="layout-element"><div class="layout layout-horizontal"><div class="layout-horizontal-elements"><div class="layout-horizontal-element"><label for="idBiology">Biology</label><div class="layout-value"></div></div><div class="layout-horizontal-element"><label for="idPsychology">Psychology</label><div class="layout-value"></div></div><div class="layout-horizontal-element"><label for="idOntology">Ontology</label><div class="layout-value"></div></div></div></div></div></div></div>`
	require.Nil(t, err)
	// require.Equal(t, expected, result)
}

func TestFormViewer(t *testing.T) {

	data := mapof.Any{
		"name":  "John Connor",
		"email": "john@connor.mil",
		"age":   42,
		"human": true,
		"ology": mapof.Any{
			"biology":    "I am a biological human",
			"psychology": "I think, therefore I am",
			"ontology":   "Here's why I am a human",
		},
	}

	form := getTestForm()
	_, err := form.Viewer(data, nil)
	// expected := `<div class="layout layout-vertical"><div class="layout-elements"><div class="layout-element"><label for="idName">Name</label><div class="layout-value"></div></div><div class="layout-element"><label for="idEmail">Email</label><div class="layout-value"></div></div><div class="layout-element"><label for="idAge">Age</label><div class="layout-value"></div></div><div class="layout-element"><label for="idHuman">Human</label><div class="layout-value"></div></div><div class="layout-element"><div class="layout layout-horizontal"><div class="layout-horizontal-elements"><div class="layout-horizontal-element"><label for="idBiology">Biology</label><div class="layout-value"></div></div><div class="layout-horizontal-element"><label for="idPsychology">Psychology</label><div class="layout-value"></div></div><div class="layout-horizontal-element"><label for="idOntology">Ontology</label><div class="layout-value"></div></div></div></div></div></div></div>`

	require.Nil(t, err)
	// require.Equal(t, expected, result)
}

func TestAllPaths(t *testing.T) {

	form := getTestElement()

	paths := form.AllElements()

	require.Equal(t, 7, len(paths))
	require.Equal(t, "name", paths[0].Path)
	require.Equal(t, "email", paths[1].Path)
	require.Equal(t, "age", paths[2].Path)
	require.Equal(t, "human", paths[3].Path)
	require.Equal(t, "ology.biology", paths[4].Path)
	require.Equal(t, "ology.psychology", paths[5].Path)
	require.Equal(t, "ology.ontology", paths[6].Path)
}

func getTestForm() form.Form {
	return form.New(getTestSchema(), getTestElement())
}

func getTestElement() form.Element {

	UseAll()

	return form.Element{
		Type: "layout-vertical",
		Children: []form.Element{
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
				Children: []form.Element{
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
