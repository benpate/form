package form

// OptionProvider is an external object that
// can inject LookupCodes based on their URL.
type LookupProvider interface {
	LookupCodes(name string) []LookupCode
}
