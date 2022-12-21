package form

// ReadOnlyLookupGroup is a simple implementation of LookupGroup that returns a static list of LookupCodes.
type ReadOnlyLookupGroup []LookupCode

func NewReadOnlyLookupGroup(codes ...LookupCode) ReadOnlyLookupGroup {
	return ReadOnlyLookupGroup(codes)
}

func (group ReadOnlyLookupGroup) Get() []LookupCode {
	return group
}
