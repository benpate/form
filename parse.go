package form

import (
	"encoding/json"

	"github.com/benpate/derp"
)

// Parse attempts to convert a value into a Form.
// Currently supports map[string]any, []byte, string, and UnmarshalMaper interface.
func Parse(data any) (Element, error) {

	const location = "form.Parse"

	result := Element{}

	switch typedData := data.(type) {

	case Element:
		return typedData, nil

	case UnmarshalMaper:
		err := result.UnmarshalMap(typedData.UnmarshalMap())
		return result, err

	case map[string]any:
		err := result.UnmarshalMap(typedData)
		return result, err

	case []byte:
		if err := json.Unmarshal(typedData, &result); err != nil {
			return result, derp.Wrap(err, location, "Invalid JSON", string(typedData))
		}
		return result, nil

	case string:
		if err := json.Unmarshal([]byte(typedData), &result); err != nil {
			return result, derp.Wrap(err, location, "Invalid JSON", typedData)
		}
		return result, nil
	}

	return result, derp.Internal(location, "Cannot Parse Value: Unknown Datatype", data)
}

// MustParse guarantees that a value has been parsed into a Form, or else it panics the application.
func MustParse(data any) Element {

	result, err := Parse(data)

	if err != nil {
		panic(err)
	}

	return result
}
