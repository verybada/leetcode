package mininumCostForTickets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinCostTickets(t *testing.T) {
	tests := []struct {
		days     []int
		costs    []int
		expected int
	}{
		{
			[]int{1, 4, 6, 7, 8, 20},
			[]int{2, 7, 15},
			11,
		},
		{
			[]int{1, 4, 6, 7, 8, 20},
			[]int{7, 2, 15},
			6,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
			[]int{2, 7, 15},
			17,
		},
		{
			[]int{3, 5, 6, 8, 9, 10, 11, 12, 13, 14, 15, 16, 20, 21, 23, 25, 26, 27, 29, 30, 33, 34, 35, 36, 38, 39, 40, 42, 45, 46, 47, 48, 49, 51, 53, 54, 56, 57, 58, 59, 60, 61, 63, 64, 67, 68, 69, 70, 72, 74, 77, 78, 79, 80, 81, 82, 83, 84, 85, 88, 91, 92, 93, 96},
			[]int{3, 17, 57},
			170,
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%+v", test.days)
		t.Run(name, func(t *testing.T) {
			cost := mincostTickets(test.days, test.costs)
			require.Equal(t, test.expected, cost)
		})
	}
}
