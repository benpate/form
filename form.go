package form

import (
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
)

type Form struct {
	Schema  schema.Schema
	Element Element
}

// New returns a fully initialized Form object (with all required values)
func New(schema schema.Schema, element Element) Form {
	return Form{
		Schema:  schema,
		Element: element,
	}
}

// Viewer creates an in-place form and executes its "Viewer" method
func Viewer(schema schema.Schema, element Element, value any, lookupProvider LookupProvider) (string, error) {
	form := New(schema, element)
	return form.Viewer(value, lookupProvider)
}

// Viewer creates an in-place form and executes its "Editorr" method
func Editor(schema schema.Schema, element Element, value any, lookupProvider LookupProvider) (string, error) {
	form := New(schema, element)
	return form.Editor(value, lookupProvider)
}

/*********************************
 * Drawing Methods
 ********************************/

// DrawString() generates this form as a string
func (form *Form) Editor(value any, lookupProvider LookupProvider) (string, error) {
	builder := html.New()
	err := form.BuildEditor(value, lookupProvider, builder)
	return builder.String(), err
}

// DrawString() generates this form as a string
func (form *Form) Viewer(value any, lookupProvider LookupProvider) (string, error) {
	builder := html.New()
	err := form.BuildViewer(value, lookupProvider, builder)
	return builder.String(), err
}

// BuildEditor generates an editable view of this form
func (form *Form) BuildEditor(value any, lookupProvider LookupProvider, builder *html.Builder) error {
	return form.Element.Edit(&form.Schema, lookupProvider, value, builder)
}

// BuildViewer generates a read-only view of this form
func (form *Form) BuildViewer(value any, lookupProvider LookupProvider, builder *html.Builder) error {
	return form.Element.View(&form.Schema, lookupProvider, value, builder)
}
