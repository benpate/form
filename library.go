package form

import (
	"github.com/benpate/derp"
)

// Library stores all of the available Renderers, and can execute them on a set of data
type Library map[string]Renderer

// New returns a fully initialized Library
func New() Library {
	return Library{}
}

// Register adds a new Renderer to the form.Library
func (library Library) Register(name string, renderer Renderer) {
	library[name] = renderer
}

// Renderer retrieves a renderer function from the library
func (library Library) Renderer(name string) (Renderer, error) {

	if renderer, ok := library[name]; ok {
		return renderer, nil
	}

	return nil, derp.New(500, "form.Library.Renderer", "Undefined Renderer", name)
}
