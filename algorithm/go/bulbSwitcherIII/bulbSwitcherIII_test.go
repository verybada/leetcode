package bulbSwitcherIII

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumTimesAllBlue(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{2, 1, 3, 5, 4},
			3,
		},
		{
			[]int{3, 2, 4, 1, 5},
			2,
		},
		{
			[]int{4, 1, 2, 3},
			1,
		},
		{
			[]int{2, 1, 4, 3, 6, 5},
			3,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			6,
		},
		{
			[]int{1, 2, 3, 4, 7, 6, 5, 8},
			6,
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%+v", test.input)
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.expected, numTimesAllBlue(test.input))
		})
	}
}
