package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxIceCream(t *testing.T) {
	tests := []struct {
		costs    []int
		coins    int
		expected int
	}{
		{[]int{1, 3, 2, 4, 1}, 7, 4},
		{[]int{10, 6, 8, 7, 7, 8}, 5, 0},
		{[]int{1, 6, 3, 1, 2, 5}, 20, 6},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v", test.costs)
		t.Run(name, func(t *testing.T) {
			result := maxIceCream(test.costs, test.coins)
			require.Equal(t, result, test.expected)
		})
	}
}
