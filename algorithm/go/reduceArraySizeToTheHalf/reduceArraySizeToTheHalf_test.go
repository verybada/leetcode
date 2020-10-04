package reduceArraySizeToTheHalf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReduceArraySizeToTheHalf(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7},
			2,
		},
		{
			[]int{7, 7, 7, 7, 7, 7},
			1,
		},
		{
			[]int{1000, 1000, 3, 7},
			1,
		},
		{
			[]int{1, 9},
			1,
		},
		{
			[]int{9, 77, 63, 22, 92, 9, 14, 54, 8, 38, 18, 19, 38, 68, 58, 19},
			5,
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%+v", test.input)
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.expected, minSetSize(test.input))
		})
	}
}
