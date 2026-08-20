package widget

import (
	"strings"

	"github.com/benpate/form"
	"github.com/benpate/html"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
)

// iif returns trueValue when the condition is met, and falseValue when it is not.
func iif[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

// isRequired returns TRUE if either the schema or the form element marks this
// field as required.
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

// getElementID returns a DOM-safe ID for this element, generating one from its
// type and path when the element does not name one.
func getElementID(element *form.Element) string {

	result := element.ID
	if result == "" {
		result = element.Type + "-" + element.Path
	}

	return strings.ReplaceAll(result, ".", "-")
}

// selectedCode returns the first LookupCode whose Value matches the one provided,
// or an empty LookupCode when nothing matches.
func selectedCode(lookupCodes []form.LookupCode, value string) form.LookupCode {

	for _, lookupCode := range lookupCodes {
		if lookupCode.Value == value {
			return lookupCode
		}
	}

	return form.LookupCode{}
}

// selectedLabels returns the Label of every LookupCode whose Value appears in the
// provided list, in the order the LookupCodes were defined.
func selectedLabels(lookupCodes []form.LookupCode, values []string) []string {

	result := make([]string, 0, len(lookupCodes))

	for _, lookupCode := range lookupCodes {
		if slice.Contains(values, lookupCode.Value) {
			result = append(result, lookupCode.Label)
		}
	}

	return result
}

// drawValue writes text into the standard read-only "layout-value" wrapper.
// InnerText escapes the text, which matters because LookupCode labels reach this
// package from a LookupProvider and may hold whatever an end-user typed.
func drawValue(b *html.Builder, text string) {
	b.Div().Class("layout-value").InnerText(text).Close()
}
