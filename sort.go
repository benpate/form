package form

import "math"

func SortLookupCodeByLabel(a LookupCode, b LookupCode) bool {
	return a.Label < b.Label
}

func SortLookupCodeByGroupThenLabel(a LookupCode, b LookupCode) bool {
	if a.Group == b.Group {
		return a.Label < b.Label
	}
	return a.Group < b.Group
}

func SortLookupCodesBySelectedValues(selected []string) func(a LookupCode, b LookupCode) bool {

	indexOf := func(slice []string, value string) int {
		for index, item := range slice {
			if item == value {
				return index
			}
		}
		return math.MaxInt
	}

	return func(a LookupCode, b LookupCode) bool {

		aIndex := indexOf(selected, a.Value)
		bIndex := indexOf(selected, b.Value)

		if aIndex == bIndex {
			return false
		}

		return aIndex < bIndex
	}
}

type LookupCodeMaker interface {
	LookupCode() LookupCode
}

func AsLookupCode[T LookupCodeMaker](maker T) LookupCode {
	return maker.LookupCode()
}
