package thousandSeperator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThousandSeparator(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{
			987,
			"987",
		},
		{
			1234,
			"1.234",
		},
		{
			123456789,
			"123.456.789",
		},
		{
			0,
			"0",
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d", test.input)
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.expected, thousandSeparator(test.input))
		})
	}
}
