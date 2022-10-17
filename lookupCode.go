package form

import (
	"strings"

	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
)

// LookupCode represents a single value/label pair
// to be used in place of Enums for optional lists.
type LookupCode struct {
	Value       string `json:"value",       form:"value"       path:"value"       bson:"value"`       // Internal value of the LookupCode
	Label       string `json:"label",       form:"label"       path:"label"       bson:"label"`       // Human-friendly label/name of the LookupCode
	Description string `json:"description", form:"description" path:"description" bson:"description"` // Optional long description of the LookupCode
	Icon        string `json:"icon",        form:"icon"        path:"icon"        bson:"icon"`        // Optional icon to use when displaying the LookupCode
	Group       string `json:"group",       form:"group"       path:"group"       bson:"group"`       // Optiional grouping to use when displaying the LookupCode
}

// NewLookupCode creates a new LookupCode from a string
func NewLookupCode(value string) LookupCode {
	return LookupCode{
		Value: value,
		Label: value,
	}
}

// ID returns the unique ID of the LookupCode, allowing them to
// be used as a set.Value
func (lookupCode LookupCode) ID() string {
	return lookupCode.Value
}

// GetLookupCodes returns a list of LookupCodes derived from:
// 1) an "enum" (string or slice-of-lookupCode) in the form element,
// 2) a "datasource" value that is looked up in the lookupProvider
// 3) a value enumerated in the schema
func GetLookupCodes(element *Element, schemaElement schema.Element, lookupProvider LookupProvider) []LookupCode {

	// If we already have an array of LookupCodes, then just return it.
	if values, ok := element.Options["enum"].(string); ok {
		if values != "" {
			return slice.Map(strings.Split(values, ","), NewLookupCode)
		}
	}

	// If we already have an array of LookupCodes, then just return it.
	if values, ok := element.Options["enum"].([]LookupCode); ok {
		if len(values) > 0 {
			return values
		}
	}

	// If we have a valid LookupProvider, then try to use it to generate lookup codes
	if lookupProvider != nil {
		if provider, ok := element.Options["provider"].(string); ok {
			return lookupProvider.LookupCodes(provider)
		}
	}

	// If we have a schemaElement (type definition), then try to use it to generate lookup codes
	if schemaElement != nil {
		enum := getSchemaEnumeration(schemaElement)
		return slice.Map(enum, NewLookupCode)
	}

	// Fall through to "no options available"
	return make([]LookupCode, 0)
}

func getSchemaEnumeration(schemaElement schema.Element) []string {

	switch s := schemaElement.(type) {

	case schema.Array:
		return getSchemaEnumeration(s.Items)
	case schema.Integer:
		return convert.SliceOfString(s.Enum)
	case schema.Number:
		return convert.SliceOfString(s.Enum)
	case schema.String:
		return s.Enum
	}

	return make([]string, 0)
}
