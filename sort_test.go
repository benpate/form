package form

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortByLabel(t *testing.T) {

	data := []LookupCode{
		{Value: "F", Label: "F"},
		{Value: "E", Label: "E"},
		{Value: "D", Label: "D"},
		{Value: "C", Label: "C"},
		{Value: "B", Label: "B"},
		{Value: "A", Label: "A"},
	}

	sort.Slice(data, func(a int, b int) bool {
		return SortLookupCodeByLabel(data[a], data[b])
	})

	require.Equal(t, "A", data[0].Label)
	require.Equal(t, "B", data[1].Label)
	require.Equal(t, "C", data[2].Label)
	require.Equal(t, "D", data[3].Label)
	require.Equal(t, "E", data[4].Label)
	require.Equal(t, "F", data[5].Label)

	sort.Slice(data, func(a int, b int) bool {
		return SortLookupCodeByLabel(data[a], data[b])
	})

	require.Equal(t, "A", data[0].Label)
	require.Equal(t, "B", data[1].Label)
	require.Equal(t, "C", data[2].Label)
	require.Equal(t, "D", data[3].Label)
	require.Equal(t, "E", data[4].Label)
	require.Equal(t, "F", data[5].Label)
}

func TestSortByBGroupThenLabel(t *testing.T) {

	data := []LookupCode{
		{Value: "F", Label: "F", Group: "A"},
		{Value: "E", Label: "E", Group: "A"},
		{Value: "D", Label: "D", Group: "A"},
		{Value: "C", Label: "C", Group: "B"},
		{Value: "B", Label: "B", Group: "B"},
		{Value: "A", Label: "A", Group: "B"},
	}

	sort.Slice(data, func(a int, b int) bool {
		return SortLookupCodeByGroupThenLabel(data[a], data[b])
	})

	require.Equal(t, "D", data[0].Label)
	require.Equal(t, "E", data[1].Label)
	require.Equal(t, "F", data[2].Label)
	require.Equal(t, "A", data[3].Label)
	require.Equal(t, "B", data[4].Label)
	require.Equal(t, "C", data[5].Label)

	sort.Slice(data, func(a int, b int) bool {
		return SortLookupCodeByGroupThenLabel(data[a], data[b])
	})

	require.Equal(t, "D", data[0].Label)
	require.Equal(t, "E", data[1].Label)
	require.Equal(t, "F", data[2].Label)
	require.Equal(t, "A", data[3].Label)
	require.Equal(t, "B", data[4].Label)
	require.Equal(t, "C", data[5].Label)
}
