package form

// registry is the system-wide registry of all form widgets.
// it must be populated with all available widgets before the system can be used.
var registry map[string]Widget

func init() {
	registry = make(map[string]Widget)
}

// Use adds a new widget into the widget registry.
func Use(name string, widget Widget) {
	registry[name] = widget
}
