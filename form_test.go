package form

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPaths(t *testing.T) {

	form := getTestForm()

	paths := form.AllPaths()

	assert.Equal(t, 5, len(paths))
	assert.Equal(t, "data.firstName", paths[0].Path)
	assert.Equal(t, "data.lastName", paths[1].Path)
	assert.Equal(t, "data.biography", paths[2].Path)
	assert.Equal(t, "data.psychology", paths[3].Path)
	assert.Equal(t, "data.ontology", paths[4].Path)
}

func getTestForm() Form {

	return Form{
		Kind: "layout-vertical",
		Children: []Form{
			{
				Kind: "text",
				Path: "data.firstName",
			}, {
				Kind: "text",
				Path: "data.lastName",
			}, {
				Kind: "layout-horizontal",
				Children: []Form{
					{
						Kind: "textarea",
						Path: "data.biography",
					}, {
						Kind: "textarea",
						Path: "data.psychology",
					}, {
						Kind: "textarea",
						Path: "data.ontology",
					},
				},
			},
		},
	}
}
