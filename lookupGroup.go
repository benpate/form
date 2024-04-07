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
