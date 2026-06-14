package groupie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	g := New()
	require.NotNil(t, g)
	require.Nil(t, g.lastValue)
}

func TestGroupie_Header(t *testing.T) {

	g := New()

	// The first distinct value starts a new group
	require.True(t, g.Header("a"))

	// Repeated values do not start a new group
	require.False(t, g.Header("a"))
	require.False(t, g.Header("a"))

	// A new value starts a new group again
	require.True(t, g.Header("b"))
	require.False(t, g.Header("b"))

	// Returning to a previous value still counts as a new group
	require.True(t, g.Header("a"))
}

func TestGroupie_Header_Sequence(t *testing.T) {

	g := New()
	values := []string{"x", "x", "y", "y", "y", "z", "x"}
	expected := []bool{true, false, true, false, false, true, true}

	for i, value := range values {
		require.Equal(t, expected[i], g.Header(value), "index %d (%q)", i, value)
	}
}

func TestGroupie_Header_Nil(t *testing.T) {

	g := New()

	// The initial lastValue is nil, so the first nil value is NOT a new group
	require.False(t, g.Header(nil))

	// A non-nil value after nil starts a group
	require.True(t, g.Header("a"))

	// Returning to nil starts a group
	require.True(t, g.Header(nil))
}

func TestGroupie_Header_MixedTypes(t *testing.T) {

	g := New()

	require.True(t, g.Header(1))
	require.False(t, g.Header(1))
	require.True(t, g.Header("1")) // different type, different value
	require.True(t, g.Header(2))
}
