package queriesOnAPermutationWithKey

import (
    "testing"

    "github.com/stretchr/testify/require"
)

func TestProcessQueries(t *testing.T) {
    tests := []struct {
        queries []int
        m int
        expectedResult []int
    } {
        {
            []int{4,1,2,2},
            4,
            []int{3,1,2,0},
        }
    }

    for _, test := range tests {
        t.Run("xx", func(t *testing.T) {
            result := processQueries(test.queries, test.m)
            require.Equal(t, test.expectedResult, result)
        })
    }
}
