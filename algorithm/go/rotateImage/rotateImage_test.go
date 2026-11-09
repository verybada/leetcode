package rotateimage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			name:     "single element",
			matrix:   [][]int{{7}},
			expected: [][]int{{7}},
		},
		{
			name:     "two by two",
			matrix:   [][]int{{1, 2}, {3, 4}},
			expected: [][]int{{3, 1}, {4, 2}},
		},
		{
			name: "three by three",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "four by four",
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{
			name: "negative numbers and zero",
			matrix: [][]int{
				{-1, 0, 2},
				{3, -4, 5},
				{0, 7, -8},
			},
			expected: [][]int{
				{0, 3, -1},
				{7, -4, 0},
				{-8, 5, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanics(t, func() {
				rotate(tt.matrix)
			})
			require.Equal(t, tt.expected, tt.matrix)
		})
	}
}
