package form

import (
	"strconv"

	"github.com/benpate/rosetta/convert"
)

// FormatValue returns the string a form writes into HTML for a value. Floats keep two decimals
// when two decimals are exact, and every digit when two decimals would round them.
func FormatValue(value any) string {

	// convert.String rounds floats to two decimals, and reports when that loses digits
	result, ok := convert.StringOk(value, "")

	if ok {
		return result
	}

	// Two decimals would round this number, and a rounded number posts back as a different
	// one, so keep every digit instead
	switch typed := value.(type) {

	case float64:
		return strconv.FormatFloat(typed, 'f', -1, 64)

	case float32:
		return strconv.FormatFloat(float64(typed), 'f', -1, 32)

	case convert.Floater:
		return strconv.FormatFloat(typed.Float(), 'f', -1, 64)
	}

	// Everything else converts as convert.String does
	return result
}

// formatValues returns the strings a form writes into HTML for a list of values, keeping every
// digit of each number, as FormatValue does.
func formatValues(value any) []string {

	// A lossless conversion needs no help
	result, ok := convert.SliceOfStringOk(value)

	if ok {
		return result
	}

	// Otherwise, format each item separately so that no number is rounded
	items, ok := convert.SliceOfAnyOk(value)

	if !ok || (len(items) != len(result)) {
		return result
	}

	formatted := make([]string, len(items))

	for index, item := range items {
		formatted[index] = FormatValue(item)
	}

	return formatted
}
