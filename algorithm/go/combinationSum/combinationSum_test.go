package combinationsum

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name: "example one", candidates: []int{2, 3, 6, 7}, target: 7,
			want: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "example two", candidates: []int{2, 3, 5}, target: 8,
			want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name: "target below all candidates", candidates: []int{2}, target: 1,
			want: [][]int{},
		},
		{
			name: "single exact match", candidates: []int{7}, target: 7,
			want: [][]int{{7}},
		},
		{
			name: "reuse single candidate", candidates: []int{2}, target: 8,
			want: [][]int{{2, 2, 2, 2}},
		},
		{
			name: "no exact sum", candidates: []int{4, 6}, target: 7,
			want: [][]int{},
		},
		{
			name: "unsorted candidates", candidates: []int{7, 3, 2, 6}, target: 7,
			want: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "skip candidates above target", candidates: []int{9, 2, 10, 3}, target: 5,
			want: [][]int{{2, 3}},
		},
		{
			name: "different orders form one combination", candidates: []int{2, 3}, target: 5,
			want: [][]int{{2, 3}},
		},
		{
			name: "backtrack after dead end", candidates: []int{3, 4}, target: 8,
			want: [][]int{{4, 4}},
		},
		{
			name: "multiple repeated combinations", candidates: []int{2, 4, 6}, target: 8,
			want: [][]int{{2, 2, 2, 2}, {2, 2, 4}, {2, 6}, {4, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.candidates, tt.target)
			for _, combination := range got {
				sort.Ints(combination)
			}
			// ElementsMatch ignores result order but still checks duplicate counts.
			require.ElementsMatch(t, tt.want, got)
		})
	}
}
