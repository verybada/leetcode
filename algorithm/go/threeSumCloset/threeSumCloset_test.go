package threesumcloset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThreeSumCloset(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected int
	}{
		{[]int{0, 0, 0}, 1, 0},
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{0, 0, 1}, 10000, 1},
		{[]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2, -2},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			require.Equal(t, test.expected, threeSumClosest(test.input, test.target))
		})
	}
}
