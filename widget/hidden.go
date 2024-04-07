package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
)

type Hidden struct{}

func (widget Hidden) View(_ *form.Element, _ *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {
	return nil
}

// Hidden registers a text <input> widget into the library
func (widget Hidden) Edit(element *form.Element, s *schema.Schema, _ form.LookupProvider, value any, b *html.Builder) error {

	if element.ReadOnly {
		return Hidden{}.View(element, s, nil, value, b)
	}

	// find the path and schema to use
	var elementValue string

	if optionValue, ok := element.Options["value"]; ok {
		elementValue = convert.String(optionValue)
	} else {
		elementValue = element.GetString(value, s)
	}

	// Start building a new tag
	b.Input("hidden", element.Path).
		ID(element.Path).
		Value(elementValue).
		Close()

	return nil
}

/***********************************
 * Wiget Metadata
 ***********************************/

func (widget Hidden) ShowLabels() bool {
	return false
}

func (widget Hidden) Encoding(_ *form.Element) string {
	return ""
}
