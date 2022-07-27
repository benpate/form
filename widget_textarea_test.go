package form

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextarea(t *testing.T) {

	form := Element{
		Type: "textarea",
		Path: "username",
	}

	html, err := form.HTML(nil, getTestSchema(), testLookupProvider{})

	assert.Nil(t, err)
	assert.Equal(t, `<textarea name="username" id="textarea-username" minlength="10" maxlength="100" pattern="[a-z]+" required="true" tabIndex="0"></textarea>`, html)
}
