package carpooling

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCarPooling(t *testing.T) {
	tests := []struct {
		trips    [][]int
		capacity int
		expected bool
	}{
		{
			[][]int{
				[]int{2, 1, 5},
				[]int{3, 3, 7},
			},
			4,
			false,
		},
		{
			[][]int{
				[]int{2, 1, 5},
				[]int{3, 3, 7},
			},
			5,
			true,
		},
		{
			[][]int{
				[]int{5, 1, 2},
			},
			4,
			false,
		},
		{
			[][]int{
				[]int{5, 1, 2},
				[]int{5, 2, 3},
				[]int{5, 3, 4},
			},
			5,
			true,
		},
		{
			[][]int{
				[]int{3, 1, 10},
				[]int{1, 2, 4},
				[]int{1, 3, 9},
			},
			3,
			false,
		},
		{
			[][]int{
				[]int{2, 2, 6},
				[]int{2, 4, 7},
				[]int{8, 6, 7},
			},
			11,
			true,
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v", test.trips)
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.expected, carPooling(test.trips, test.capacity))
		})
	}
}
