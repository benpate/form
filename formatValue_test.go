package form

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestFormatValue confirms that floats keep two decimals when they are exact and every digit when
// they are not, and that every other value converts as convert.String does
func TestFormatValue(t *testing.T) {

	tests := []struct {
		value    any
		expected string
	}{
		{1.0, "1"},
		{1.5, "1.50"},
		{2.25, "2.25"},
		{-0.5, "-0.50"},
		{0.125, "0.125"},
		{1.005, "1.005"},
		{40.712776, "40.712776"},
		{-74.005974, "-74.005974"},
		{0.001, "0.001"},
		{float32(0.125), "0.125"},
		{float32(0.1), "0.10"},
		{7, "7"},
		{"0.125", "0.125"},
		{true, "true"},
		{nil, ""},
	}

	for _, test := range tests {
		require.Equal(t, test.expected, FormatValue(test.value), "value: %#v", test.value)
	}
}

// TestFormatValue_RoundTrip confirms that every formatted float parses back to the number it came from
func TestFormatValue_RoundTrip(t *testing.T) {

	for _, value := range []float64{0.5, 1.5, 2.25, 0.125, 1.005, 0.1, 0.3, 1.0 / 3.0, 40.712776, -74.005974, 1e-7, 123456.789} {
		parsed, err := strconv.ParseFloat(FormatValue(value), 64)
		require.NoError(t, err, "value: %v", value)
		require.Equal(t, value, parsed, "value: %v", value)
	}
}

// TestFormatValues confirms that every number in a list keeps the digits it needs
func TestFormatValues(t *testing.T) {

	require.Equal(t, []string{"1.50", "0.125"}, formatValues([]any{1.5, 0.125}))
	require.Equal(t, []string{"1.50", "0.125"}, formatValues([]float64{1.5, 0.125}))
	require.Equal(t, []string{"a", "b"}, formatValues([]string{"a", "b"}))
	require.Equal(t, []string{"0.125"}, formatValues(0.125))
	require.Equal(t, []string{}, formatValues(nil))
}
