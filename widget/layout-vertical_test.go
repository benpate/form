package widget

import (
	"encoding/json"
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/mapof"
	"github.com/stretchr/testify/require"
)

func TestLayoutVertical(t *testing.T) {

	UseAll()

	element := form.Element{
		Type:  "layout-vertical",
		Label: "This is my Vertical Layout",
		Children: []form.Element{
			{
				Type:  "text",
				Label: "Name",
				Path:  "name",
			},
			{
				Type:  "text",
				Label: "Email",
				Path:  "email",
			},
			{
				Type:  "text",
				Label: "Age",
				Path:  "age",
			},
		},
	}

	value := mapof.Any{
		"name":  "John Connor",
		"email": "john@resistance.mil",
		"age":   27,
	}

	schema := getTestSchema()
	builder := html.New()
	err := element.Edit(&schema, nil, &value, builder)

	require.Nil(t, err)
	// expected := `<div class="layout layout-vertical"><div class="layout-title">This is my Vertical Layout</div><div class="layout-vertical-elements"><div class="layout-vertical-element"><label>Name</label><input name="name" id="text-name" value="John Connor" type="text" maxlength="50" tabIndex="0"></div><div class="layout-vertical-element"><label>Email</label><input name="email" id="text-email" value="john@resistance.mil" type="email" minlength="10" maxlength="100" required="true" tabIndex="0"></div><div class="layout-vertical-element"><label>Age</label><input name="age" id="text-age" value="27" type="number" step="1" min="10" max="100" required="true" tabIndex="0"></div></div></div>`
	// require.Equal(t, expected, builder.String())
}

func TestRules(t *testing.T) {

	form := form.Element{
		Type:  "layout-vertical",
		Label: "This is my Vertical Layout",
		Children: []form.Element{
			{
				Type:  "text",
				Label: "Name",
				Path:  "name",
			},
			{
				Type:  "text",
				Label: "Email",
				Path:  "email",
			},
			{
				Type:  "text",
				Label: "Age",
				Path:  "age",
			},
		},
	}

	builder := html.New()
	schema := getTestSchema()
	err := form.Edit(&schema, nil, nil, builder)
	require.Nil(t, err)

	// expected := `<div class="layout layout-vertical"><div class="layout-title">This is my Vertical Layout</div><div class="layout-vertical-elements"><div class="layout-vertical-element"><label>Name</label><input name="name" id="text-name" type="text" maxlength="50" tabIndex="0"></div><div class="layout-vertical-element"><label>Email</label><input name="email" id="text-email" type="email" minlength="10" maxlength="100" required="true" tabIndex="0"></div><div class="layout-vertical-element"><label>Age</label><input name="age" id="text-age" type="number" step="1" min="10" max="100" required="true" tabIndex="0"></div></div></div>`
	// require.Equal(t, expected, builder.String())
}

func TestLayoutVertical_Unmarshal(t *testing.T) {

	var element form.Element

	formJSON := `{
		"type": "layout-tabs",
		"label":"Edit Follow Settings",
		"id":"tab-container",
		"children":[{
			"type":"layout-vertical",
			"id":"tab-settings",
			"label":"Settings",
			"children": [
				{
					"type":"text",
					"label":"Fediverse Address or Website URL",
					"path":"url",
					"description":"Enter the URL of the website you want to subscribe to."
				},
				{
					"type":"select",
					"label":"Folder",
					"path":"folderId",
					"options":{"provider": "folders"},
					"description": "Automatically add items to this folder."
				}
			]
		}, {
			"type":"layout-vertical",
			"id":"tab-info",
			"label":"Info",
			"readOnly":true,
			"children":[{
				"type":"text",
				"label":"Method",
				"path":"method"
			}, {
				"type":"text",
				"label":"Status",
				"path":"status"
			}, {
				"type":"text",
				"label":"Notes",
				"path":"statusMessag"
			}]
		}]
	}`

	require.Nil(t, json.Unmarshal([]byte(formJSON), &element))

	form := form.New(getTestSchema(), element)

	_, err := form.Editor(nil, nil)

	// expected := `<div class="layout-title">Edit Follow Settings</div><div class="tabs" data-script="install TabContainer"><div role="tablist"><button type="button" role="tab" id="tab-tab-settings" class="tab-label" aria-controls="panel-tab-settings" tabIndex="0" aria-selected="true">Settings</button><button type="button" role="tab" id="tab-tab-info" class="tab-label" aria-controls="panel-tab-info" tabIndex="0">Info</button></div><div role="tabpanel" id="panel-tab-settings" aria-labelledby="tab-tab-settings"><div class="layout layout-vertical"><div class="layout-vertical-elements"><div class="layout-vertical-element"><label>Fediverse Address or Website URL</label><input name="url" id="text-url" type="text" tabIndex="0"><div class="text-sm gray40">Enter the URL of the website you want to subscribe to.</div></div><div class="layout-vertical-element"><label>Folder</label><select id="select-folderId" name="folderId" tabIndex="0"></select><div class="text-sm gray40">Automatically add items to this folder.</div></div></div></div></div><div role="tabpanel" id="panel-tab-info" aria-labelledby="tab-tab-info" hidden="true"><div class="layout layout-vertical"><div class="layout-vertical-elements"><div class="layout-vertical-element"><label>Method</label><div class="layout-value"></div></div><div class="layout-vertical-element"><label>Status</label><div class="layout-value"></div></div><div class="layout-vertical-element"><label>Notes</label><div class="layout-value"></div></div></div></div></div></div>`

	require.Nil(t, err)
}
