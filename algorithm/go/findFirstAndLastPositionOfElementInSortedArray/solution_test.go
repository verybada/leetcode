package findFirstAndLastPositionOfElementInSortedArray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "nil input", target: 0, want: []int{-1, -1}},
		{name: "empty input", nums: []int{}, target: 0, want: []int{-1, -1}},
		{name: "example found", nums: []int{5, 7, 7, 8, 8, 10}, target: 8, want: []int{3, 4}},
		{name: "example missing", nums: []int{5, 7, 7, 8, 8, 10}, target: 6, want: []int{-1, -1}},
		{name: "single element found", nums: []int{5}, target: 5, want: []int{0, 0}},
		{name: "single element missing", nums: []int{5}, target: 0, want: []int{-1, -1}},
		{name: "all elements match", nums: []int{2, 2, 2, 2}, target: 2, want: []int{0, 3}},
		{name: "duplicates at start", nums: []int{1, 1, 2, 3}, target: 1, want: []int{0, 1}},
		{name: "duplicates at end", nums: []int{1, 2, 3, 3}, target: 3, want: []int{2, 3}},
		{name: "unique middle match", nums: []int{1, 3, 5}, target: 3, want: []int{1, 1}},
		{name: "first element", nums: []int{1, 3, 5}, target: 1, want: []int{0, 0}},
		{name: "last element", nums: []int{1, 3, 5}, target: 5, want: []int{2, 2}},
		{name: "target below minimum", nums: []int{1, 3, 5}, target: -1, want: []int{-1, -1}},
		{name: "target above maximum", nums: []int{1, 3, 5}, target: 7, want: []int{-1, -1}},
		{name: "negative target", nums: []int{-5, -3, -3, -1, 0}, target: -3, want: []int{1, 2}},
		{name: "zero target", nums: []int{-1, 0, 0, 1}, target: 0, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, searchRange(tt.nums, tt.target))
		})
	}
}
