package widget

import "github.com/benpate/form"

type testLookupProvider struct{}

func (t testLookupProvider) Group(_ string) form.LookupGroup {
	return form.NewReadOnlyLookupGroup(
		form.LookupCode{
			Label: "This is the first code",
			Value: "ONE",
		},
		form.LookupCode{
			Label: "This is the second code",
			Value: "TWO",
		},
		form.LookupCode{
			Label: "This is the third code",
			Value: "THREE",
		},
		form.LookupCode{
			Label: "This is the fourth code",
			Value: "FOUR",
		},
		form.LookupCode{
			Label: "This is the fifth code",
			Value: "FIVE",
		},
	)
}
