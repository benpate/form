package form

import (
	"testing"

	"github.com/benpate/rosetta/maps"
	"github.com/stretchr/testify/require"
)

func TestMultiselect(t *testing.T) {

	form := Element{
		Type: "multiselect",
		Path: "tags",
	}

	value := maps.Map{"tags": []string{"pretty", "please"}}

	{
		html, err := form.HTML(value, getTestSchema(), nil)
		require.Nil(t, err)
		require.Equal(t, `<div class="multiselect" data-script="install multiselect(sort:false)"><div class="options" style="maxHeight:300px"><label for="tags_pretty"><input type="checkbox" name="tags" id="tags_pretty" value="pretty" checked="true"><div><div>pretty</div></div></label><label for="tags_please"><input type="checkbox" name="tags" id="tags_please" value="please" checked="true"><div><div>please</div></div></label><label for="tags_my"><input type="checkbox" name="tags" id="tags_my" value="my"><div><div>my</div></div></label><label for="tags_dear"><input type="checkbox" name="tags" id="tags_dear" value="dear"><div><div>dear</div></div></label><label for="tags_aunt"><input type="checkbox" name="tags" id="tags_aunt" value="aunt"><div><div>aunt</div></div></label><label for="tags_sally"><input type="checkbox" name="tags" id="tags_sally" value="sally"><div><div>sally</div></div></label></div></div>`, html)
	}

	{
		form.Options = maps.Map{"sort": true}

		html, err := form.HTML(value, getTestSchema(), nil)
		require.Nil(t, err)
		require.Equal(t, `<div class="multiselect" data-script="install multiselect(sort:true)"><div class="options" style="maxHeight:300px"><label for="tags_pretty"><input type="checkbox" name="tags" id="tags_pretty" value="pretty" checked="true"><div><div>pretty</div></div></label><label for="tags_please"><input type="checkbox" name="tags" id="tags_please" value="please" checked="true"><div><div>please</div></div></label><label for="tags_my"><input type="checkbox" name="tags" id="tags_my" value="my"><div><div>my</div></div></label><label for="tags_dear"><input type="checkbox" name="tags" id="tags_dear" value="dear"><div><div>dear</div></div></label><label for="tags_aunt"><input type="checkbox" name="tags" id="tags_aunt" value="aunt"><div><div>aunt</div></div></label><label for="tags_sally"><input type="checkbox" name="tags" id="tags_sally" value="sally"><div><div>sally</div></div></label></div><div class="buttons"><button type="button" data-sort="up">△</button><button type="button" data-sort="down">▽</button></div></div>`, html)
	}

	// t.Log(html)
}
