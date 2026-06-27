package form

// FilterLookupCodeByGroup returns a predicate that matches LookupCodes
// belonging to the named group.
func FilterLookupCodeByGroup(group string) func(LookupCode) bool {
	return func(code LookupCode) bool {
		return code.Group == group
	}
}
