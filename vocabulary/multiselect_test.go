package vocabulary

import (
	"testing"

	"github.com/benpate/form"
	"github.com/stretchr/testify/require"
)

func TestMultiselect(t *testing.T) {

	library := getTestLibrary()
	s := getTestSchema()

	f := form.Form{
		Kind: "multiselect",
		Path: "tags",
	}

	value := map[string]interface{}{
		"tags": []string{"pretty", "please"},
	}

	{
		html, err := f.HTML(&library, s, value)
		require.Nil(t, err)
		require.Equal(t, `<div class="multiselect" data-script="install multiselect(sort:false)"><div class="options" style="maxHeight:300px"><label for="tags_pretty"><input type="checkbox" name="tags" id="tags_pretty" value="pretty" checked="true"><div><div>pretty</div></div></label><label for="tags_please"><input type="checkbox" name="tags" id="tags_please" value="please" checked="true"><div><div>please</div></div></label><label for="tags_my"><input type="checkbox" name="tags" id="tags_my" value="my"><div><div>my</div></div></label><label for="tags_dear"><input type="checkbox" name="tags" id="tags_dear" value="dear"><div><div>dear</div></div></label><label for="tags_aunt"><input type="checkbox" name="tags" id="tags_aunt" value="aunt"><div><div>aunt</div></div></label><label for="tags_sally"><input type="checkbox" name="tags" id="tags_sally" value="sally"><div><div>sally</div></div></label></div></div>`, html)
	}

	{
		f.Options = map[string]any{"sort": true}

		html, err := f.HTML(&library, s, value)
		require.Nil(t, err)
		require.Equal(t, `<div class="multiselect" data-script="install multiselect(sort:true)"><div class="options" style="maxHeight:300px"><label for="tags_pretty"><input type="checkbox" name="tags" id="tags_pretty" value="pretty" checked="true"><div><div>pretty</div></div></label><label for="tags_please"><input type="checkbox" name="tags" id="tags_please" value="please" checked="true"><div><div>please</div></div></label><label for="tags_my"><input type="checkbox" name="tags" id="tags_my" value="my"><div><div>my</div></div></label><label for="tags_dear"><input type="checkbox" name="tags" id="tags_dear" value="dear"><div><div>dear</div></div></label><label for="tags_aunt"><input type="checkbox" name="tags" id="tags_aunt" value="aunt"><div><div>aunt</div></div></label><label for="tags_sally"><input type="checkbox" name="tags" id="tags_sally" value="sally"><div><div>sally</div></div></label></div><div class="buttons"><button type="button" data-sort="up">△</button><button type="button" data-sort="down">▽</button></div></div>`, html)
	}

	// t.Log(html)
}
