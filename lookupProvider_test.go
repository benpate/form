package form

type testLookupProvider struct{}

func (t testLookupProvider) Group(_ string) LookupGroup {
	return NewReadOnlyLookupGroup(
		LookupCode{
			Label: "This is the first code",
			Value: "ONE",
		},
		LookupCode{
			Label: "This is the second code",
			Value: "TWO",
		},
		LookupCode{
			Label: "This is the third code",
			Value: "THREE",
		},
		LookupCode{
			Label: "This is the fourth code",
			Value: "FOUR",
		},
		LookupCode{
			Label: "This is the fifth code",
			Value: "FIVE",
		},
	)
}
