package form

// OptionProvider is an external object that
// can inject LookupCodes based on their URL.
type LookupProvider interface {
	Group(name string) LookupGroup
}

type LookupGroup interface {
	Get() []LookupCode
}

type WritableLookupGroup interface {
	LookupGroup
	Add(name string) (string, error)
}
