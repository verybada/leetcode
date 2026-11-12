package addbinary

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{
			name:     "both zero",
			a:        "0",
			b:        "0",
			expected: "0",
		},
		{
			name:     "zero on left",
			a:        "0",
			b:        "101",
			expected: "101",
		},
		{
			name:     "zero on right",
			a:        "101",
			b:        "0",
			expected: "101",
		},
		{
			name:     "single digit carry",
			a:        "1",
			b:        "1",
			expected: "10",
		},
		{
			name:     "addition without carry",
			a:        "1010",
			b:        "101",
			expected: "1111",
		},
		{
			name:     "longer left operand",
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			name:     "longer right operand",
			a:        "1",
			b:        "11",
			expected: "100",
		},
		{
			name:     "carry within result",
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
		{
			name:     "carry stops before highest digit",
			a:        "1011",
			b:        "1",
			expected: "1100",
		},
		{
			name:     "carry with both bits set",
			a:        "111",
			b:        "111",
			expected: "1110",
		},
		{
			name:     "carry beyond machine integer size",
			a:        strings.Repeat("1", 100),
			b:        "1",
			expected: "1" + strings.Repeat("0", 100),
		},
		{
			name:     "maximum length left operand with full carry",
			a:        strings.Repeat("1", 10000),
			b:        "1",
			expected: "1" + strings.Repeat("0", 10000),
		},
		{
			name:     "maximum length right operand with full carry",
			a:        "1",
			b:        strings.Repeat("1", 10000),
			expected: "1" + strings.Repeat("0", 10000),
		},
		{
			name:     "both operands at maximum length",
			a:        strings.Repeat("1", 10000),
			b:        strings.Repeat("1", 10000),
			expected: strings.Repeat("1", 10000) + "0",
		},
		{
			name:     "maximum length without carry",
			a:        "1" + strings.Repeat("0", 9999),
			b:        "1",
			expected: "1" + strings.Repeat("0", 9998) + "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, addBinary(tt.a, tt.b))
		})
	}
}
