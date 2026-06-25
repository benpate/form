package form

import (
	"github.com/benpate/derp"
	"github.com/benpate/exp"
	"github.com/benpate/rosetta/schema"
)

// Validate returns an error if the form references any field that does not exist
// in the schema.
func (form *Form) Validate() error {

	const location = "form.Form.Validate"

	// Call this once, when a form is loaded or saved, to catch definition mistakes
	// (such as a misspelled "show-if" field) before an end-user submits data.
	if err := form.Element.validateSchema(&form.Schema); err != nil {
		return derp.Wrap(err, location, "Invalid form definition")
	}

	return nil
}

// validateSchema confirms that this element (and all of its children) only
// reference fields that are defined in the schema.
func (element *Element) validateSchema(s *schema.Schema) error {

	const location = "form.Element.validateSchema"

	// RULE: If the element has a Path, then that Path must exist in the schema.
	if element.Path != "" {
		if _, ok := s.GetElement(element.Path); !ok {
			return derp.Internal(location, "Form references a path that is not in the schema", element.Path)
		}
	}

	// RULE: Every field referenced by a "show-if" expression must exist in the schema.
	if showIf := element.Options.GetString("show-if"); showIf != "" {
		for _, field := range exp.Parse(showIf).Fields() {
			if _, ok := s.GetElement(field); !ok {
				return derp.Internal(location, "Form 'show-if' references a field that is not in the schema", field, showIf)
			}
		}
	}

	// Recursively validate all child elements.
	for index := range element.Children {
		if err := element.Children[index].validateSchema(s); err != nil {
			return derp.Wrap(err, location, "Invalid child element", element.Children[index].Path)
		}
	}

	return nil
}
