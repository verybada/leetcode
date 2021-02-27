package equalsubstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEqualSubstring(t *testing.T) {
	tests := []struct {
		source   string
		target   string
		maxCost  int
		expected int
	}{
		{"abcd", "bcdf", 3, 3},
		{"abcd", "cdef", 3, 1},
		{"abcd", "acde", 0, 1},
		{"cdef", "abcd", 3, 1},
		{"krrgw", "zjxss", 19, 2},
		{"krpgjbjjznpzdfy", "nxargkbydxmsgby", 14, 4},
	}

	for _, test := range tests {
		t.Run(test.source, func(t *testing.T) {
			require.Equal(t, test.expected,
				equalSubstring(test.source, test.target, test.maxCost))
		})
	}
}
