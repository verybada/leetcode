package permuteii

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "single element",
			nums:     []int{7},
			expected: [][]int{{7}},
		},
		{
			name:     "all elements equal",
			nums:     []int{2, 2, 2},
			expected: [][]int{{2, 2, 2}},
		},
		{
			name:     "repeated elements",
			nums:     []int{1, 1, 2},
			expected: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			name:     "nonadjacent repeated elements",
			nums:     []int{1, 2, 1},
			expected: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			name: "distinct elements",
			nums: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name: "two repeated pairs",
			nums: []int{1, 2, 1, 2},
			expected: [][]int{
				{1, 1, 2, 2},
				{1, 2, 1, 2},
				{1, 2, 2, 1},
				{2, 1, 1, 2},
				{2, 1, 2, 1},
				{2, 2, 1, 1},
			},
		},
		{
			name:     "negative numbers and zero",
			nums:     []int{-1, 0, -1},
			expected: [][]int{{-1, -1, 0}, {-1, 0, -1}, {0, -1, -1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.ElementsMatch(t, tt.expected, permuteUnique(tt.nums))
		})
	}
}
