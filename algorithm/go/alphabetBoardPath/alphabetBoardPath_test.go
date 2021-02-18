package alphabetBoardPath

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAlphabetBoardPath(t *testing.T) {
	tests := []struct {
		target string
		path   string
	}{
		{
			"leet",
			"DDR!UURRR!!DDD!",
		},
		{
			"code",
			"RR!DDRR!UUL!R!",
		},
		{
			"zdz",
			"DDDDD!UUUUURRR!DDDDLLLD!",
		},
	}

	for _, test := range tests {
		t.Run(test.target, func(t *testing.T) {
			require.Equal(t, test.path, alphabetBoardPath(test.target))
		})
	}
}
