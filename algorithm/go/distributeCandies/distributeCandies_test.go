package distributecandies

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDistributeCandies(t *testing.T) {
	tests := []struct {
		candies   int
		numPeople int
		expected  []int
	}{
		{7, 4, []int{1, 2, 3, 1}},
		{10, 3, []int{5, 2, 3}},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d_%d", test.candies, test.numPeople)
		t.Run(name, func(t *testing.T) {
			require.Equal(t,
				test.expected,
				distributeCandies(test.candies, test.numPeople))
		})
	}
}

func TestDistributeCandiesGoroutine(t *testing.T) {
	tests := []struct {
		candies   int
		numPeople int
		expected  []int
	}{
		{7, 4, []int{1, 2, 3, 1}},
		{10, 3, []int{5, 2, 3}},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d_%d", test.candies, test.numPeople)
		t.Run(name, func(t *testing.T) {
			require.Equal(t,
				test.expected,
				distributeCandiesGoroutine(test.candies, test.numPeople))
		})
	}
}
