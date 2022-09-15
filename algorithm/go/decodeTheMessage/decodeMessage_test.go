package decodeMessage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeMessage(t *testing.T) {
	tests := []struct {
		key      string
		message  string
		expected string
	}{
		{
			"the quick brown fox jumps over the lazy dog",
			"vkbs bs t suepuv",
			"this is a secret",
		},
		{
			"eljuxhpwnyrdgtqkviszcfmabo",
			"zwx hnfx lqantp mnoeius ycgk vcnjrdb",
			"the five boxing wizards jump quickly",
		},
	}

	for _, test := range tests {
		t.Run(test.key, func(t *testing.T) {
			result := decodeMessage(test.key, test.message)
			require.Equal(t, test.expected, result)
		})
	}
}
