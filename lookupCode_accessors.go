package form

import "github.com/benpate/rosetta/schema"

func LookupCodeSchema() schema.Element {
	return schema.Object{
		Properties: schema.ElementMap{
			"value":       schema.String{},
			"label":       schema.String{},
			"description": schema.String{},
			"icon":        schema.String{},
			"group":       schema.String{},
			"href":        schema.String{},
		},
	}
}

func (lookupCode *LookupCode) GetPointer(name string) (any, bool) {

	switch name {

	case "value":
		return &lookupCode.Value, true

	case "label":
		return &lookupCode.Label, true

	case "description":
		return &lookupCode.Description, true

	case "icon":
		return &lookupCode.Icon, true

	case "group":
		return &lookupCode.Group, true

	case "href":
		return &lookupCode.Href, true

	}

	return nil, false
}
