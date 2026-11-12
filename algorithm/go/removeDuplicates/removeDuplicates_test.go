package removeduplicates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "first example",
			input:    []int{1, 1, 1, 2, 2, 3},
			expected: []int{1, 1, 2, 2, 3},
		},
		{
			name:     "second example",
			input:    []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expected: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "two equal elements",
			input:    []int{1, 1},
			expected: []int{1, 1},
		},
		{
			name:     "two distinct elements",
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "all equal",
			input:    []int{5, 5, 5, 5, 5},
			expected: []int{5, 5},
		},
		{
			name:     "all distinct",
			input:    []int{-2, -1, 0, 1, 2},
			expected: []int{-2, -1, 0, 1, 2},
		},
		{
			name:     "every value occurs twice",
			input:    []int{-1, -1, 0, 0, 1, 1},
			expected: []int{-1, -1, 0, 0, 1, 1},
		},
		{
			name:     "excess duplicates at beginning",
			input:    []int{1, 1, 1, 1, 2, 3},
			expected: []int{1, 1, 2, 3},
		},
		{
			name:     "excess duplicates at end",
			input:    []int{1, 2, 3, 3, 3, 3},
			expected: []int{1, 2, 3, 3},
		},
		{
			name:     "consecutive oversized groups",
			input:    []int{-3, -3, -3, -2, -2, -2, 0, 0, 0, 1},
			expected: []int{-3, -3, -2, -2, 0, 0, 1},
		},
		{
			name:     "boundary values",
			input:    []int{-10000, -10000, -10000, 0, 10000, 10000, 10000},
			expected: []int{-10000, -10000, 0, 10000, 10000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeDuplicates(tt.input)
			require.Equal(t, len(tt.expected), k)
			require.Equal(t, tt.expected, tt.input[:k])
		})
	}
}
