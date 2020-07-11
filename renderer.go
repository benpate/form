package form

import (
	"strings"

	"github.com/benpate/schema"
)

// Renderer is a function signature that writes HTML for a fully populated widget into a string builder.
type Renderer func(Form, schema.Schema, interface{}, *strings.Builder) error
