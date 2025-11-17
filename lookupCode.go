package form

import (
	"strings"

	"github.com/benpate/rosetta/convert"
	"github.com/benpate/rosetta/mapof"
	"github.com/benpate/rosetta/schema"
	"github.com/benpate/rosetta/slice"
	"github.com/benpate/rosetta/sliceof"
)

// LookupCode represents a single value/label pair
// to be used in place of Enums for optional lists.
type LookupCode struct {
	Value       string `json:"value,omitempty"       form:"value"       bson:"value,omitempty"`       // Internal value of the LookupCode
	Label       string `json:"label,omitempty"       form:"label"       bson:"label,omitempty"`       // Human-friendly label/name of the LookupCode
	Description string `json:"description,omitempty" form:"description" bson:"description,omitempty"` // Optional long description of the LookupCode
	Icon        string `json:"icon,omitempty"        form:"icon"        bson:"icon,omitempty"`        // Optional icon to use when displaying the LookupCode
	Group       string `json:"group,omitempty"       form:"group"       bson:"group,omitempty"`       // Optiional grouping to use when displaying the LookupCode
	Href        string `json:"href,omitempty"        form:"href"        bson:"href,omitempty"`        // Optional URL to use when using this LookupCode
}

// NewLookupCode creates a new LookupCode from a string
func NewLookupCode(value string) LookupCode {
	return LookupCode{
		Value: value,
		Label: value,
	}
}

// ParseLookupCode converts a number of values into a LookupCode, including:
// LookupCode, , mapof.Any, mapof.String, map[string]any, map[string]string, and
// string. For strings, the string value is used for both the .Value and .Label
// If no compatible type is found, then an empty LookupCode is returned.
func ParseLookupCode(value any) LookupCode {

	switch typed := value.(type) {

	case LookupCode:
		return typed

	case string:

		return LookupCode{
			Value: typed,
			Label: typed,
		}

	case mapof.Any:

		return LookupCode{
			Value:       typed.GetString("value"),
			Label:       typed.GetString("label"),
			Description: typed.GetString("description"),
			Icon:        typed.GetString("icon"),
			Group:       typed.GetString("group"),
			Href:        typed.GetString("href"),
		}

	case mapof.String:

		return LookupCode{
			Value:       typed.GetString("value"),
			Label:       typed.GetString("label"),
			Description: typed.GetString("description"),
			Icon:        typed.GetString("icon"),
			Group:       typed.GetString("group"),
			Href:        typed.GetString("href"),
		}

	case map[string]any:
		return ParseLookupCode(mapof.Any(typed))

	case map[string]string:
		return ParseLookupCode(mapof.String(typed))
	}

	return LookupCode{}
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
//
// The boolean value is TRUE if this comes from a WritableLookupGroup
func GetLookupCodes(element *Element, schemaElement schema.Element, lookupProvider LookupProvider) ([]LookupCode, bool) {

	// If we have a valid LookupProvider, then try to use it to generate lookup codes first
	if lookupProvider != nil {
		if provider, ok := element.Options["provider"].(string); ok {
			group := lookupProvider.Group(provider)

			_, isWritable := group.(WritableLookupGroup)
			return group.Get(), isWritable
		}
	}

	// If an "enum" option is present, then try to use it to generate LookupCodes
	if enumValue, ok := element.Options["enum"]; ok {

		switch typed := enumValue.(type) {

		case string:
			return slice.Map(strings.Split(typed, ","), NewLookupCode), false

		case sliceof.Object[LookupCode]:
			return typed, false

		case []LookupCode:
			return typed, false

		case []string:
			return slice.Map(typed, NewLookupCode), false

		case []any:
			return slice.Map(typed, ParseLookupCode), false
		}
	}

	// Last, if we have a schemaElement (type definition), then try to use it to generate lookup codes
	if schemaElement != nil {
		enum := getSchemaEnumeration(schemaElement)
		return slice.Map(enum, NewLookupCode), false
	}

	// Fall through to "no options available"
	return make([]LookupCode, 0), false
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

	return make([]string, 0) // This should probably never happen.
}
