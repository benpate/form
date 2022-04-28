package form

// OptionProvider is an external object that
// can inject OptionCodes based on their URL.
type OptionProvider interface {
	OptionCodes(string) ([]OptionCode, error)
}

// OptionCode represents a single value/label pair
// to be used in place of Enums for optional lists.
type OptionCode struct {
	Value       string // Internal value of the Option
	Label       string // Human-friendly label/name of the Option
	Description string // Optional long description of the Option
	Icon        string // Optional icon to use when displaying the Option
	Group       string // Optiional grouping to use when displaying the Option
}
