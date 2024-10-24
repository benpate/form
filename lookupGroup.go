package form

// ReadOnlyLookupGroup is a simple implementation of the
// LookupGroup interface that returns a static list of LookupCodes.
type ReadOnlyLookupGroup []LookupCode

// NewReadOnlyLookupGroup returns a fully initialized ReadOnlyLookupGroup
func NewReadOnlyLookupGroup(codes ...LookupCode) ReadOnlyLookupGroup {
	return ReadOnlyLookupGroup(codes)
}

// Get returns the list of LookupCodes
func (group ReadOnlyLookupGroup) Get() []LookupCode {
	return group
}

// Value returns the LookupCode that matches the provided value,
// and an empty LookupCode if no match is found.
func (group ReadOnlyLookupGroup) Value(value string) LookupCode {
	for _, code := range group {
		if code.Value == value {
			return code
		}
	}
	return LookupCode{}
}
