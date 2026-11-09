package permute

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPermute(t *testing.T) {
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
			name:     "two elements",
			nums:     []int{0, 1},
			expected: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "three elements",
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
			name: "negative numbers and zero",
			nums: []int{-1, 0, 2},
			expected: [][]int{
				{-1, 0, 2},
				{-1, 2, 0},
				{0, -1, 2},
				{0, 2, -1},
				{2, -1, 0},
				{2, 0, -1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.ElementsMatch(t, tt.expected, permute(tt.nums))
		})
	}
}
