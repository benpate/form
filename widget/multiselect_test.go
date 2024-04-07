package widget

import (
	"testing"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/sliceof"
	"github.com/stretchr/testify/require"
)

func TestMultiselect(t *testing.T) {

	element := Element{
		Type: "multiselect",
		Path: "tags",
	}

	schema := getTestSchema()
	value := mapof.Any{"tags": sliceof.String{"pretty", "please"}}

	{
		builder := html.New()
		err := element.Edit(&schema, nil, &value, builder)
		require.Nil(t, err)
		require.Equal(t, `<div class="multiselect" data-script="install multiselect(sort:false)"><div class="options" style="maxHeight:300px"><label for="multiselect-tags-pretty"><input type="checkbox" name="tags" id="multiselect-tags-pretty" value="pretty" checked="true"><div><div>pretty</div></div></label><label for="multiselect-tags-please"><input type="checkbox" name="tags" id="multiselect-tags-please" value="please" checked="true"><div><div>please</div></div></label><label for="multiselect-tags-my"><input type="checkbox" name="tags" id="multiselect-tags-my" value="my"><div><div>my</div></div></label><label for="multiselect-tags-dear"><input type="checkbox" name="tags" id="multiselect-tags-dear" value="dear"><div><div>dear</div></div></label><label for="multiselect-tags-aunt"><input type="checkbox" name="tags" id="multiselect-tags-aunt" value="aunt"><div><div>aunt</div></div></label><label for="multiselect-tags-sally"><input type="checkbox" name="tags" id="multiselect-tags-sally" value="sally"><div><div>sally</div></div></label></div></div>`, builder.String())
	}

	{
		element.Options = mapof.Any{"sort": true}

		builder := html.New()
		err := element.Edit(&schema, nil, &value, builder)
		require.Nil(t, err)
		require.Equal(t, `<div class="multiselect" data-script="install multiselect(sort:true)"><div class="options" style="maxHeight:300px"><label for="multiselect-tags-pretty"><input type="checkbox" name="tags" id="multiselect-tags-pretty" value="pretty" checked="true"><div><div>pretty</div></div></label><label for="multiselect-tags-please"><input type="checkbox" name="tags" id="multiselect-tags-please" value="please" checked="true"><div><div>please</div></div></label><label for="multiselect-tags-my"><input type="checkbox" name="tags" id="multiselect-tags-my" value="my"><div><div>my</div></div></label><label for="multiselect-tags-dear"><input type="checkbox" name="tags" id="multiselect-tags-dear" value="dear"><div><div>dear</div></div></label><label for="multiselect-tags-aunt"><input type="checkbox" name="tags" id="multiselect-tags-aunt" value="aunt"><div><div>aunt</div></div></label><label for="multiselect-tags-sally"><input type="checkbox" name="tags" id="multiselect-tags-sally" value="sally"><div><div>sally</div></div></label></div><div class="buttons"><button type="button" data-sort="up">△</button><button type="button" data-sort="down">▽</button></div></div>`, builder.String())
	}

}
