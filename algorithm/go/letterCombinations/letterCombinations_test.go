package letterCombinations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"2", []string{"a", "b", "c"}},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			require.Equal(t, test.expected, letterCombinations(test.input))
		})
	}
}
