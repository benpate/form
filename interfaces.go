package form

import (
	"github.com/benpate/html"
)

// LookupProvider is an external object that
// can inject LookupCodes based on their URL.
type LookupProvider interface {
	Group(name string) LookupGroup
}

// LookupGroup is an read-only interface that returns a list of LookupCodes
type LookupGroup interface {
	Get() []LookupCode
}

// WriteableLookupGroup is a read-write interface that returns
// a list of LookupCodes, and can add new codes to the list.
type WritableLookupGroup interface {
	LookupGroup
	Add(name string) (string, error)
}

// Widget defines a data type that can be included in a form
type Widget interface {
	View(form *Form, element *Element, lookupProvider LookupProvider, value any, builder *html.Builder) error
	Edit(form *Form, element *Element, lookupProvider LookupProvider, value any, builder *html.Builder) error
	ShowLabels() bool
	Encoding(element *Element) string
}

// UnmarshalMaper wraps the UnmarshalMap interface
type UnmarshalMaper interface {

	// UnmarshalMap returns a value in the format map[string]interface
	UnmarshalMap() map[string]any
}
