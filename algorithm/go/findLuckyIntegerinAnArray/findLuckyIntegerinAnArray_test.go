package findLuckyIntegerinAnArray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindLucky(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{2, 2, 3, 4},
			2,
		},
		{
			[]int{1, 2, 2, 3, 3, 3},
			3,
		},
		{
			[]int{2, 2, 2, 3, 3},
			-1,
		},
		{
			[]int{5},
			-1,
		},
		{
			[]int{7, 7, 7, 7, 7, 7, 7},
			7,
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%+v", test.input)
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.expected, findLucky(test.input))
		})
	}
}
