package form

import (
	"testing"

	"github.com/benpate/derp"
	"github.com/benpate/html"
	"github.com/benpate/schema"
	"github.com/stretchr/testify/assert"
)

func TestLibraryWidget(t *testing.T) {

	library := New(nil)

	library.Register("test", func(form Form, _ *schema.Schema, _ interface{}, b *html.Builder) error {
		b.Empty("SAMPLE-WIDGET")
		return nil
	})

	form := Form{
		Kind: "test",
	}

	html, err := form.HTML(library, &schema.Schema{}, nil)

	assert.Equal(t, "<SAMPLE-WIDGET>", html)
	assert.Nil(t, err)
}

func TestLibraryError(t *testing.T) {

	library := New(nil)

	library.Register("error", func(form Form, _ *schema.Schema, _ interface{}, b *html.Builder) error {
		return derp.New(500, "Error", "error")
	})

	form := Form{Kind: "error"}

	html, err := form.HTML(library, &schema.Schema{}, nil)

	assert.Equal(t, "", html)
	assert.NotNil(t, err)
}

func TestLibraryNotFound(t *testing.T) {

	library := New(nil)

	library.Register("test", func(form Form, _ *schema.Schema, _ interface{}, b *html.Builder) error {
		b.Empty("SAMPLE-WIDGET")
		return nil
	})

	form := Form{Kind: "not-found"}

	html, err := form.HTML(library, &schema.Schema{}, nil)

	assert.Equal(t, "", html)
	assert.NotNil(t, err)
}
