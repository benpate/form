package vocabulary

import "github.com/benpate/form"

func getTestLibrary() form.Library {

	library := form.New()

	All(library)

	return library
}
