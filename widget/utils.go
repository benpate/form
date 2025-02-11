package widget

import (
	"github.com/benpate/form"
	"github.com/benpate/rosetta/schema"
)

func iif[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

func isRequired(element *form.Element, schemaElement schema.Element) bool {

	if schemaElement != nil {
		if schemaElement.IsRequired() {
			return true
		}
	}

	if element != nil {
		if element.Options.GetBool("required") {
			return true
		}
	}

	return false
}
