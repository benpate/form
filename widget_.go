package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

// registry is the system-wide registry of all form widgets
var registry map[string]WidgetFunc

func init() {
	registry = make(map[string]WidgetFunc)
}

// WidgetFunc is a function signature that writes HTML for a fully populated widget into a string builder.
type WidgetFunc func(element *Element, schema *schema.Schema, lookupProvider LookupProvider, value any, builder *html.Builder) error

// TODO:
// date
// datetime
// time

// Register adds a new widget into the widget registry.
func Register(name string, widget WidgetFunc) {
	registry[name] = widget
}
