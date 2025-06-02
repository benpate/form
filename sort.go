package form

// SortLookupCodeByLabel is a sort function that works with the sort.Slice
// function.
func SortLookupCodeByLabel(a LookupCode, b LookupCode) int {

	if a.Label < b.Label {
		return -1
	}

	if a.Label > b.Label {
		return 1
	}

	return 0
}

// SortLookupCodeByGroupThenLabel is a sort function that works with the
// sort.Slice function.
func SortLookupCodeByGroupThenLabel(a LookupCode, b LookupCode) int {

	if a.Group < b.Group {
		return -1
	}

	if a.Group > b.Group {
		return 1
	}

	if a.Label < b.Label {
		return -1
	}

	if a.Label > b.Label {
		return 1
	}

	return 0
}

// LookupCodeMaker is an interface that wraps the LookupCode method
type LookupCodeMaker interface {
	// LookupCode returns the data from current object in the form of a form.LookupCode
	LookupCode() LookupCode
}

// AsLookupCode is a helper function that converts any object that implements
// the LookupCodeMaker interface into a form.LookupCode
func AsLookupCode[T LookupCodeMaker](maker T) LookupCode {
	return maker.LookupCode()
}
