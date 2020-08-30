package generateParenthesis

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n        int
		expected []string
	}{
		{1, []string{"()"}},
		{2, []string{"(())", "()()"}},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d", test.n)
		t.Run(name, func(t *testing.T) {

			require.Equal(t, test.expected, generateParenthesis(test.n))
		})
	}
}
