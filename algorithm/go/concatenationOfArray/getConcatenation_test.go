package concatenationOfArray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{1, 2, 3}, []int{1, 2, 3, 1, 2, 3},
		},
	}

	for _, test := range tests {
		result := getConcatenation(test.nums)
		require.Equal(t, result, test.expected)
	}
}
