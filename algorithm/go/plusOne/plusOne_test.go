package plusone

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "increment without carry",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "zero",
			input:    []int{0},
			expected: []int{1},
		},
		{
			name:     "single digit without carry",
			input:    []int{8},
			expected: []int{9},
		},
		{
			name:     "carry to previous digit",
			input:    []int{1, 2, 9},
			expected: []int{1, 3, 0},
		},
		{
			name:     "carry across multiple digits",
			input:    []int{1, 9, 9},
			expected: []int{2, 0, 0},
		},
		{
			name:     "single digit overflow",
			input:    []int{9},
			expected: []int{1, 0},
		},
		{
			name:     "all digits overflow",
			input:    []int{9, 9, 9},
			expected: []int{1, 0, 0, 0},
		},
		{
			name:     "internal zeroes",
			input:    []int{1, 0, 0, 9},
			expected: []int{1, 0, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, plusOne(tt.input))
		})
	}
}
