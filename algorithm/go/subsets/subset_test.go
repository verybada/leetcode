package subset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubSets(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				[]int{},
				[]int{1},
				[]int{2},
				[]int{1, 2},
				[]int{3},
				[]int{1, 3},
				[]int{2, 3},
				[]int{1, 2, 3},
			},
		},
		//{
		//[]int{9, 0, 3, 5, 7},
		//[][]int{
		//[]int{},
		//[]int{9},
		//[]int{0},
		//[]int{0, 9},
		//[]int{3},
		//[]int{3, 9},
		//[]int{0, 3},
		//[]int{0, 3, 9},
		//[]int{5},
		//[]int{5, 9},
		//[]int{0, 5},
		//[]int{0, 5, 9},
		//[]int{3, 5},
		//[]int{3, 5, 9},
		//[]int{0, 3, 5},
		//[]int{0, 3, 5, 9},
		//[]int{7},
		//[]int{7, 9},
		//[]int{0, 7},
		//[]int{0, 7, 9},
		//[]int{3, 7},
		//[]int{3, 7, 9},
		//[]int{0, 3, 7},
		//[]int{0, 3, 7, 9},
		//[]int{5, 7},
		//[]int{5, 7, 9},
		//[]int{0, 5, 7},
		//[]int{0, 5, 7, 9},
		//[]int{3, 5, 7},
		//[]int{3, 5, 7, 9},
		//[]int{0, 3, 5, 7},
		//[]int{0, 3, 5, 7, 9},
		//},
		//},
	}
	for _, test := range tests {
		name := fmt.Sprintf("%+v", test.input)
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.expected, subsets(test.input))
		})
	}
}
