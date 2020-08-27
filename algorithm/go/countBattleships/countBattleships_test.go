package countBattleships

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountBattleships(t *testing.T) {
	tests := []struct {
		board [][]byte
		count int
	}{
		{
			[][]byte{
				[]byte{'X', '.', '.', 'X'},
				[]byte{'.', '.', '.', 'X'},
				[]byte{'.', '.', '.', 'X'},
			},
			2,
		},
		{
			[][]byte{
				[]byte{'X', '.', '.', 'X'},
				[]byte{'.', 'X', '.', 'X'},
				[]byte{'.', 'X', '.', 'X'},
			},
			3,
		},
		{
			[][]byte{
				[]byte{'X', '.', 'X'},
				[]byte{'.', 'X', '.'},
				[]byte{'X', '.', 'X'},
			},
			5,
		},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%s", test.board)
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.count, countBattleships(test.board))
		})
	}
}
