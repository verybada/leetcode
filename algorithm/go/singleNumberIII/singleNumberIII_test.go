package singlenumberiii

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleNumberII(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
		{[]int{-1, 0}, []int{-1, 0}},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v", test.nums)
		t.Run(name, func(t *testing.T) {
			results := singleNumberIII(test.nums)
			sort.Ints(results)
			require.Equal(t, test.expected, results)
		})
	}
}
