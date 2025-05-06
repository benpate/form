package form

import (
	"encoding/json"

	"github.com/benpate/derp"
)

// Parse attempts to convert a value into a Form.
// Currently supports map[string]any, []byte, string, and UnmarshalMaper interface.
func Parse(data any) (Element, error) {

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
		err := json.Unmarshal(typedData, &result)
		return result, err

	case string:
		err := json.Unmarshal([]byte(typedData), &result)
		return result, err
	}

	return result, derp.InternalError("form.Parse", "Cannot Parse Value: Unknown Datatype", data)
}

// MustParse guarantees that a value has been parsed into a Form, or else it panics the application.
func MustParse(data any) Element {

	result, err := Parse(data)

	if err != nil {
		panic(err)
	}

	return result
}
