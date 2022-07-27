package form

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPaths(t *testing.T) {

	form := getTestForm()

	paths := form.AllElements()

	assert.Equal(t, 5, len(paths))
	assert.Equal(t, "data.firstName", paths[0].Path)
	assert.Equal(t, "data.lastName", paths[1].Path)
	assert.Equal(t, "data.biography", paths[2].Path)
	assert.Equal(t, "data.psychology", paths[3].Path)
	assert.Equal(t, "data.ontology", paths[4].Path)
}

func getTestForm() Element {

	return Element{
		Type: "layout-vertical",
		Children: []Element{
			{
				Type: "text",
				Path: "data.firstName",
			}, {
				Type: "text",
				Path: "data.lastName",
			}, {
				Type: "layout-horizontal",
				Children: []Element{
					{
						Type: "textarea",
						Path: "data.biography",
					}, {
						Type: "textarea",
						Path: "data.psychology",
					}, {
						Type: "textarea",
						Path: "data.ontology",
					},
				},
			},
		},
	}
}
