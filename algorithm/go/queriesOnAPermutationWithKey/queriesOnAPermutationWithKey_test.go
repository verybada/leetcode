package queriesOnAPermutationWithKey

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProcessQueries(t *testing.T) {
	tests := []struct {
		queries        []int
		m              int
		expectedResult []int
	}{
		{
			[]int{3, 1, 2, 1},
			5,
			[]int{2, 1, 2, 1},
		},
		{
			[]int{4, 1, 2, 2},
			4,
			[]int{3, 1, 2, 0},
		},
		{
			[]int{7, 5, 5, 8, 3},
			8,
			[]int{6, 5, 0, 7, 5},
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%+v", test.queries)
		t.Run(name, func(t *testing.T) {
			result := processQueries(test.queries, test.m)
			require.Equal(t, test.expectedResult, result)
		})
	}
}
