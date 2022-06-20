package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

// Renderer is a function signature that writes HTML for a fully populated widget into a string builder.
type Renderer func(Form, *schema.Schema, any, *html.Builder) error
