package vocabulary

import (
	"testing"

	"github.com/benpate/form"
	"github.com/stretchr/testify/assert"
)

func getTestLibrary() form.Library {

	library := form.New()

	All(library)

	return library
}

func TestAll(t *testing.T) {

	library := getTestLibrary()

	assert.NotNil(t, library)
}
