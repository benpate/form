package vocabulary

import "github.com/benpate/form"

type testOptionProvider bool

func (t testOptionProvider) OptionCodes(_ string) ([]form.OptionCode, error) {
	return []form.OptionCode{
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
	}, nil
}
