package form

import (
	"encoding/json"

	"github.com/benpate/derp"
)

// UnmarshalJSON implements the json.Unmarshaler interface for Form, validating
// the decoded form against its schema before returning.
func (form *Form) UnmarshalJSON(data []byte) error {

	const location = "form.Form.UnmarshalJSON"

	// formAlias drops Form's UnmarshalJSON method to avoid infinite recursion.
	type formAlias Form

	var decoded formAlias

	if err := json.Unmarshal(data, &decoded); err != nil {
		return derp.Wrap(err, location, "Invalid JSON", string(data))
	}

	*form = Form(decoded)

	if err := form.Validate(); err != nil {
		return derp.Wrap(err, location, "Form is not valid")
	}

	return nil
}
