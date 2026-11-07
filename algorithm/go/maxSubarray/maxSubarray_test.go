package maxsubarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "mixed example", nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		{name: "single positive", nums: []int{1}, want: 1},
		{name: "whole array example", nums: []int{5, 4, -1, 7, 8}, want: 23},
		{name: "single negative", nums: []int{-5}, want: -5},
		{name: "single zero", nums: []int{0}, want: 0},
		{name: "all negative", nums: []int{-8, -3, -6, -2, -5, -4}, want: -2},
		{name: "all positive", nums: []int{1, 2, 3, 4}, want: 10},
		{name: "all zero", nums: []int{0, 0, 0}, want: 0},
		{name: "zero among negatives", nums: []int{-2, 0, -1}, want: 0},
		{name: "best prefix", nums: []int{4, 3, -10, 1, 2}, want: 7},
		{name: "best suffix", nums: []int{-5, 1, -3, 4, 2}, want: 6},
		{name: "isolated maximum", nums: []int{-3, 5, -8}, want: 5},
		{name: "bridge a negative value", nums: []int{4, -1, 2, 1}, want: 6},
		{name: "restart after large loss", nums: []int{2, -10, 3, 4}, want: 7},
		{name: "equal best sums", nums: []int{2, -3, 2}, want: 2},
		{name: "must remain contiguous", nums: []int{5, -10, 6}, want: 6},
		{name: "large sum", nums: []int{10000, 10000, -1, 10000}, want: 29999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, maxSubArray(tt.nums))
		})
	}
}
