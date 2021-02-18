package reformatDate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReformatDate(t *testing.T) {
	tests := []struct {
		date     string
		expected string
	}{
		{
			"20th Oct 2052",
			"2052-10-20",
		},
		{
			"6th Jun 1933",
			"1933-06-06",
		},
		{
			"26th May 1960",
			"1960-05-26",
		},
	}
	for _, test := range tests {
		t.Run(test.date, func(t *testing.T) {
			require.Equal(t, test.expected, reformatDate(test.date))
		})
	}
}
