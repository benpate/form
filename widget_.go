package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

type Widget interface {
	View(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, builder *html.Builder) error
	Edit(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, builder *html.Builder) error
	ShowLabels() bool
}

// registry is the system-wide registry of all form widgets
var registry map[string]Widget

func init() {
	registry = make(map[string]Widget)
}

// Register adds a new widget into the widget registry.
func Register(name string, widget Widget) {
	registry[name] = widget
}
