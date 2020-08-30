package frequencySort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFrequencySort(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"tree", "eetr"},
		{"cccaaa", "cccaaa"},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			require.Equal(t, test.expected, frequencySort(test.input))
		})
	}
}
