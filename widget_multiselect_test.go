package form

import (
	"testing"

	"github.com/benpate/html"
	"github.com/benpate/rosetta/maps"
	"github.com/stretchr/testify/require"
)

func TestMultiselect(t *testing.T) {

	element := Element{
		Type: "multiselect",
		Path: "tags",
	}

	schema := getTestSchema()
	value := maps.Map{"tags": []string{"pretty", "please"}}

	{
		builder := html.New()
		err := element.Edit(&schema, nil, value, builder)
		require.Nil(t, err)
		require.Equal(t, `<div class="multiselect" data-script="install multiselect(sort:false)"><div class="options" style="maxHeight:300px"><label for="tags.pretty"><input type="checkbox" name="tags" id="tags.pretty" value="pretty" checked="true"><div><div>pretty</div></div></label><label for="tags.please"><input type="checkbox" name="tags" id="tags.please" value="please" checked="true"><div><div>please</div></div></label><label for="tags.my"><input type="checkbox" name="tags" id="tags.my" value="my"><div><div>my</div></div></label><label for="tags.dear"><input type="checkbox" name="tags" id="tags.dear" value="dear"><div><div>dear</div></div></label><label for="tags.aunt"><input type="checkbox" name="tags" id="tags.aunt" value="aunt"><div><div>aunt</div></div></label><label for="tags.sally"><input type="checkbox" name="tags" id="tags.sally" value="sally"><div><div>sally</div></div></label></div></div>`, builder.String())
	}

	{
		element.Options = maps.Map{"sort": true}

		builder := html.New()
		err := element.Edit(&schema, nil, value, builder)
		require.Nil(t, err)
		require.Equal(t, `<div class="multiselect" data-script="install multiselect(sort:true)"><div class="options" style="maxHeight:300px"><label for="tags.pretty"><input type="checkbox" name="tags" id="tags.pretty" value="pretty" checked="true"><div><div>pretty</div></div></label><label for="tags.please"><input type="checkbox" name="tags" id="tags.please" value="please" checked="true"><div><div>please</div></div></label><label for="tags.my"><input type="checkbox" name="tags" id="tags.my" value="my"><div><div>my</div></div></label><label for="tags.dear"><input type="checkbox" name="tags" id="tags.dear" value="dear"><div><div>dear</div></div></label><label for="tags.aunt"><input type="checkbox" name="tags" id="tags.aunt" value="aunt"><div><div>aunt</div></div></label><label for="tags.sally"><input type="checkbox" name="tags" id="tags.sally" value="sally"><div><div>sally</div></div></label></div><div class="buttons"><button type="button" data-sort="up">△</button><button type="button" data-sort="down">▽</button></div></div>`, builder.String())
	}

	// t.Log(html)
}
