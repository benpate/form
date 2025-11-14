package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/convert"
)

type Hidden struct{}

func (widget Hidden) View(_ *form.Form, _ *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {
	return nil
}

// Hidden registers a text <input> widget into the library
func (widget Hidden) Edit(f *form.Form, e *form.Element, _ form.LookupProvider, value any, b *html.Builder) error {

	// find the path and schema to use
	var elementValue string

	if optionValue, ok := e.Options["value"]; ok {
		elementValue = convert.String(optionValue)
	} else {
		elementValue = e.GetString(value, &f.Schema)
	}

	// Start building a new tag
	b.Input("hidden", e.Path).
		ID(e.Path).
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
