package romantoint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		roman string
		value int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, test := range tests {
		t.Run(test.roman, func(t *testing.T) {
			require.Equal(t, test.value, romanToInt(test.roman))
		})
	}
}
