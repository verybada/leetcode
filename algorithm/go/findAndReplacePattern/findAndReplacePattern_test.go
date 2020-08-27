package findAndReplacePattern

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindAndReplacePattern(t *testing.T) {
	tests := []struct {
		words    []string
		pattern  string
		expected []string
	}{
		{
			[]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
			"abb",
			[]string{"mee", "aqq"},
		},
		{
			[]string{},
			"abcdefg",
			[]string{},
		},
		{
			[]string{"jojoabc", "jojodef", "abcdefg", "a"},
			"ababxyz",
			[]string{"jojoabc", "jojodef"},
		},
	}
	for _, test := range tests {
		t.Run(test.pattern, func(t *testing.T) {

			require.Equal(
				t, test.expected,
				findAndReplacePattern(test.words, test.pattern))
		})
	}
}
