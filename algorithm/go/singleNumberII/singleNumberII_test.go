package singlenumberii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleNumberII(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v", test.nums)
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.expected, singleNumberII(test.nums))
		})
	}
}
