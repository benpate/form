package vocabulary

import (
	"github.com/benpate/convert"
	"github.com/benpate/path"
	"github.com/benpate/schema"
)

// locateSchema looks up schema and values using a variable path.
func locateSchema(pathString string, originalSchema schema.Schema, value interface{}) (schema.Schema, string) {

	var resultSchema schema.Schema
	var resultValue string

	resultSchema = schema.Any{}

	// If the path is empty, then return empty values
	if pathString != "" {

		// Parse the path to the field value.
		pathObject := path.New(pathString)

		// If the schema is nil, then there's not much we can do here.
		if originalSchema != nil {
			if s, err := originalSchema.Path(pathObject); err == nil {
				resultSchema = s
			}
		}

		if value, err := pathObject.Get(value); err == nil {
			resultValue, _ = convert.StringOk(value, "")
		}
	}

	return resultSchema, resultValue
}
