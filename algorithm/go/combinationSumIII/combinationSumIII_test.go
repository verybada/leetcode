package combinationSumIII

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombinationSumIII(t *testing.T) {
	tests := []struct {
		k        int
		n        int
		expected [][]int
	}{
		{
			3, 7, [][]int{
				[]int{1, 2, 4},
			},
		},
		{
			3, 9, [][]int{
				[]int{1, 2, 6},
				[]int{1, 3, 5},
				[]int{2, 3, 4},
			},
		},
		{
			4, 1, [][]int{},
		},
		{
			3, 2, [][]int{},
		},
		{
			9, 45, [][]int{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			8, 36, [][]int{
				[]int{1, 2, 3, 4, 5, 6, 7, 8},
			},
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%+v", test.expected)
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.expected, combinationSum3(test.k, test.n))
		})
	}
}
