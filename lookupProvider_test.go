package form

type testLookupProvider struct{}

func (t testLookupProvider) LookupCodes(_ string) []LookupCode {
	return []LookupCode{
		{
			Label: "This is the first code",
			Value: "ONE",
		},
		{
			Label: "This is the second code",
			Value: "TWO",
		},
		{
			Label: "This is the third code",
			Value: "THREE",
		},
		{
			Label: "This is the fourth code",
			Value: "FOUR",
		},
		{
			Label: "This is the fifth code",
			Value: "FIVE",
		},
	}
}
