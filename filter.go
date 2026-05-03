package form

func FilterLookupCodeByGroup(group string) func(LookupCode) bool {
	return func(code LookupCode) bool {
		return code.Group == group
	}
}
