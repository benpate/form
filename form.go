package form

import (
	"github.com/benpate/derp"
	"github.com/benpate/html"
	"github.com/benpate/schema"
)

type Form struct {
	Element Element
	Library *Library
	Schema  *schema.Schema
}

// HTML returns a populated HTML string for the provided form, schema, and value
func (form Form) HTML(library *Library, schema *schema.Schema, value any) (string, error) {

	b := html.New()

	if err := element.Write(library, schema, value, b); err != nil {
		return "", derp.Wrap(err, "form.HTML", "Error rendering element", form)
	}

	return b.String(), nil
}
