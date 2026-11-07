package combinationsumii

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name: "example one", candidates: []int{10, 1, 2, 7, 6, 1, 5}, target: 8,
			want: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			name: "example two", candidates: []int{2, 5, 2, 1, 2}, target: 5,
			want: [][]int{{1, 2, 2}, {5}},
		},
		{
			name: "single exact match", candidates: []int{3}, target: 3,
			want: [][]int{{3}},
		},
		{
			name: "cannot reuse one element", candidates: []int{2}, target: 4,
			want: [][]int{},
		},
		{
			name: "equal elements at different indices", candidates: []int{2, 2}, target: 4,
			want: [][]int{{2, 2}},
		},
		{
			name: "deduplicate repeated pairs", candidates: []int{1, 1, 1, 1}, target: 2,
			want: [][]int{{1, 1}},
		},
		{
			name: "deduplicate exact matches", candidates: []int{3, 3, 3}, target: 3,
			want: [][]int{{3}},
		},
		{
			name: "target below all candidates", candidates: []int{3, 4, 5}, target: 2,
			want: [][]int{},
		},
		{
			name: "target above total sum", candidates: []int{1, 2, 3}, target: 7,
			want: [][]int{},
		},
		{
			name: "no exact sum", candidates: []int{2, 4, 6}, target: 5,
			want: [][]int{},
		},
		{
			name: "unsorted candidates", candidates: []int{5, 3, 1, 2}, target: 6,
			want: [][]int{{1, 2, 3}, {1, 5}},
		},
		{
			name: "multiple combinations with duplicates", candidates: []int{1, 1, 2, 2, 3}, target: 4,
			want: [][]int{{1, 1, 2}, {1, 3}, {2, 2}},
		},
		{
			name: "use every element", candidates: []int{1, 2, 4}, target: 7,
			want: [][]int{{1, 2, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.candidates, tt.target)
			for _, combination := range got {
				sort.Ints(combination)
			}
			// Ignore ordering while still rejecting duplicate combinations.
			require.ElementsMatch(t, tt.want, got)
		})
	}
}
