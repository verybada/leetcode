package groupThePeople

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupThePeople(t *testing.T) {
	tests := []struct {
		groupSizes []int
		expected   [][]int
	}{
		{
			[]int{3, 3, 3, 3, 3, 1, 3},
			[][]int{{0, 1, 2}, {5}, {3, 4, 6}},
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v", test.groupSizes)
		t.Run(name, func(t *testing.T) {

			result := groupThePeople(test.groupSizes)
			require.Equal(t, test.expected, result)
		})
	}
}
