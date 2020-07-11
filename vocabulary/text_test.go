package vocabulary

import (
	"testing"

	"github.com/benpate/form"
)

func TestText(t *testing.T) {

	library := getTestLibrary()

	f := form.Form{
		Kind: "text",
		Path: "abc.xyz",
	}

	html, err := f.HTML(library, nil, nil)

	t.Log(html)
	derp.Report(err)
}
