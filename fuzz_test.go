package form

import (
	"encoding/json"
	"strings"
	"testing"
)

// seedElementJSON is a collection of byte slices used to seed the Element-parsing
// fuzzers. It mixes valid forms, malformed JSON, empty/degenerate input, and
// adversarial shapes (deep nesting, wrong types) so the fuzzer starts from
// interesting territory.
var seedElementJSON = [][]byte{
	[]byte(``),
	[]byte(`{}`),
	[]byte(`null`),
	[]byte(`[]`),
	[]byte(`"just a string"`),
	[]byte(`{not valid json`),
	[]byte(`{"type":"text","path":"name"}`),
	[]byte(`{"type":"text","path":"name","label":"Name","readOnly":true}`),
	[]byte(`{"type":"layout-vertical","children":[{"type":"text","path":"name"}]}`),
	[]byte(`{"type":"text","options":{"show-if":"a is b","enum":["x","y"]}}`),
	[]byte(`{"children":[{"children":[{"children":[{"type":"text"}]}]}]}`),
	[]byte(`{"type":123,"path":true,"readOnly":"not-a-bool"}`),
	[]byte(`{"options":"not-a-map","children":"not-an-array"}`),
	[]byte(`{"children":[null,42,"text",{"type":"text"}]}`),
}

// FuzzParse exercises Parse with arbitrary bytes (as both []byte and string
// input) and asserts that it never panics and never returns a non-empty
// Element alongside an error.
func FuzzParse(f *testing.F) {

	for _, seed := range seedElementJSON {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, data []byte) {

		// Parsing raw bytes must never panic, regardless of input.
		bytesElement, bytesErr := Parse(data)

		// RULE: an error result must come with an empty Element.
		if bytesErr != nil && !bytesElement.IsEmpty() {
			t.Fatalf("Parse([]byte) returned an error AND a non-empty element: %v", bytesErr)
		}

		// Parsing the same data as a string must agree with the []byte path.
		stringElement, stringErr := Parse(string(data))

		if (bytesErr == nil) != (stringErr == nil) {
			t.Fatalf("Parse disagreed between []byte and string for the same input")
		}

		// A successfully parsed Element must survive a re-marshal without panicking.
		if bytesErr == nil {
			if _, err := json.Marshal(bytesElement); err != nil {
				t.Fatalf("Parsed element failed to re-marshal: %v", err)
			}
		}

		_ = stringElement
	})
}

// FuzzForm_UnmarshalJSON exercises Form.UnmarshalJSON (which also runs Validate)
// with arbitrary bytes and asserts that it never panics, and that any form which
// unmarshals cleanly can be re-marshaled.
func FuzzForm_UnmarshalJSON(f *testing.F) {

	// Seed with the element shapes wrapped as a full Form, plus standalone forms.
	for _, seed := range seedElementJSON {
		f.Add(seed)
	}

	f.Add([]byte(`{"schema":{"type":"object","properties":{"name":{"type":"string"}}},"form":{"type":"text","path":"name"}}`))
	f.Add([]byte(`{"schema":{"type":"object","properties":{"name":{"type":"string"}}},"form":{"type":"text","path":"missing"}}`))
	f.Add([]byte(`{"schema":{},"form":{"type":"text","options":{"show-if":"ghost is true"}}}`))

	f.Fuzz(func(t *testing.T, data []byte) {

		var form Form

		// Unmarshalling arbitrary bytes must never panic.
		err := json.Unmarshal(data, &form)

		// A form that unmarshalled cleanly must re-marshal without panicking.
		if err == nil {
			if _, marshalErr := json.Marshal(form); marshalErr != nil {
				t.Fatalf("Unmarshalled form failed to re-marshal: %v", marshalErr)
			}

			// Validate is idempotent: a form that passed validation during
			// Unmarshal must still pass when validated again.
			if validateErr := form.Validate(); validateErr != nil {
				t.Fatalf("Form passed Unmarshal validation but failed a second Validate: %v", validateErr)
			}
		}
	})
}

// FuzzElement_UnmarshalMap exercises the recursive map-descent parser by feeding
// it maps decoded from arbitrary JSON, asserting that it never panics no matter
// how the child tree is shaped.
func FuzzElement_UnmarshalMap(f *testing.F) {

	for _, seed := range seedElementJSON {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, data []byte) {

		// Decode the fuzz bytes into a generic map. Inputs that are not JSON
		// objects are uninteresting for this target, so skip them.
		var decoded map[string]any
		if err := json.Unmarshal(data, &decoded); err != nil {
			return
		}

		// UnmarshalMap must never panic, even on deeply nested or wrongly-typed children.
		var element Element
		_ = element.UnmarshalMap(decoded)
	})
}

// FuzzMustParse confirms that MustParse only panics for inputs that Parse also
// rejects, and never panics for inputs that Parse accepts.
func FuzzMustParse(f *testing.F) {

	for _, seed := range seedElementJSON {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, data []byte) {

		// Determine whether Parse accepts this input.
		_, parseErr := Parse(data)

		// If Parse succeeds, MustParse must not panic.
		if parseErr == nil {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("MustParse panicked on input that Parse accepted: %v", r)
				}
			}()

			_ = MustParse(data)
			return
		}

		// If Parse fails, MustParse must panic. Recover so the fuzzer continues.
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("MustParse did not panic on input that Parse rejected: %q", strings.TrimSpace(string(data)))
			}
		}()

		_ = MustParse(data)
	})
}
