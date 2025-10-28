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

func (lookupCode LookupCode) GetStringOK(name string) (string, bool) {

	switch name {

	case "value":
		return lookupCode.Value, true

	case "label":
		return lookupCode.Label, true

	case "description":
		return lookupCode.Description, true

	case "icon":
		return lookupCode.Icon, true

	case "group":
		return lookupCode.Group, true

	case "href":
		return lookupCode.Href, true

	}

	return "", false
}

func (lookupCode *LookupCode) SetString(name string, value string) bool {

	switch name {

	case "value":
		lookupCode.Value = value
		return true

	case "label":
		lookupCode.Label = value
		return true

	case "description":
		lookupCode.Description = value
		return true

	case "icon":
		lookupCode.Icon = value
		return true

	case "group":
		lookupCode.Group = value
		return true

	case "href":
		lookupCode.Href = value
		return true

	}

	return false
}
