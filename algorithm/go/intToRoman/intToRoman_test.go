package inttoroman

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		value int
		roman string
	}{
		{58, "LVIII"},
		{1994, "MCMXCIV"},
		{3000, "MMM"},
		{3749, "MMMDCCXLIX"},
	}
	for _, test := range tests {
		t.Run(test.roman, func(t *testing.T) {
			require.Equal(t, test.roman, intToRoman(test.value))
		})
	}
}
