package stringMatchingInAnArray

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringMatching(t *testing.T) {
	tests := []struct {
		words    []string
		expected []string
	}{
		{
			[]string{"mass", "as", "hero", "superhero"},
			[]string{"as", "hero"},
		},
		{
			[]string{"leetcode", "et", "code"},
			[]string{"et", "code"},
		},
		{
			[]string{"blue", "green", "yellow"},
			[]string{},
		},
		{
			[]string{"leetcoder", "leetcode", "od", "hamlet", "am"},
			[]string{"leetcode", "od", "am"},
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v", test.words)
		t.Run(name, func(t *testing.T) {
			results := stringMatching(test.words)
			sort.Strings(results)
			sort.Strings(test.expected)
			require.Equal(t, test.expected, results)
		})
	}
}
