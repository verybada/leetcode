package lastStoneWeight

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{2, 7, 4, 1, 8, 1},
			1,
		},
		{
			[]int{1},
			1,
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("%+v", test.input)
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.expected, lastStoneWeight(test.input))
		})
	}
}
