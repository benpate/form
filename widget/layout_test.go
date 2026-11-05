package widget

import (
	"net/url"
	"strings"
	"testing"

	"github.com/benpate/form"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/stretchr/testify/require"
)

// TestLayout_SavesOnlyWhatItDraws checks that every widget which skips a child by its Form
// options draws, and saves, the same set of fields.
func TestLayout_SavesOnlyWhatItDraws(t *testing.T) {

	UseAll()

	s := schema.New(schema.Object{Properties: schema.ElementMap{
		"name":   schema.String{},
		"bridge": schema.Boolean{},
		"guest":  schema.String{},
	}})

	test := func(layoutType string, shown bool, options ...string) {
		t.Helper()

		f := form.New(s, form.Element{Type: layoutType, Children: []form.Element{
			{Type: "text", Path: "name"},
			{Type: "toggle", Path: "bridge", Options: mapof.Template{"show-if-option": "show-bridge"}},
			{Type: "text", Path: "guest", Options: mapof.Template{"hide-if-option": "is-member"}},
		}}, options...)

		object := mapof.Any{"name": "Sarah", "bridge": false, "guest": "stored"}

		// The editor draws both optional fields, or neither...
		html, err := f.Editor(object, nil)
		require.NoError(t, err)
		require.Equal(t, shown, strings.Contains(html, `name="bridge"`), "%s %v", layoutType, options)
		require.Equal(t, shown, strings.Contains(html, `name="guest"`), "%s %v", layoutType, options)

		// ...and a post that names them both writes exactly the fields it drew
		values := url.Values{"name": {"Sam"}, "bridge": {"true"}, "guest": {"posted"}}
		require.NoError(t, f.SetURLValues(&object, values, nil))
		require.Equal(t, "Sam", object["name"], "%s %v", layoutType, options)
		require.Equal(t, shown, object["bridge"], "%s %v", layoutType, options)
		require.Equal(t, shown, object["guest"] == "posted", "%s %v", layoutType, options)
	}

	for _, layoutType := range []string{"layout-vertical", "layout-horizontal", "layout-group", "container"} {
		test(layoutType, false, "show-bridge:false", "is-member:true")
		test(layoutType, true, "show-bridge:true", "is-member:false")
	}
}
