package defusethebomb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecrypt(t *testing.T) {
	tests := []struct {
		code     []int
		k        int
		expected []int
	}{
		{[]int{5, 7, 1, 4}, 3, []int{12, 10, 16, 13}},
		{[]int{1, 2, 3, 4}, 0, []int{0, 0, 0, 0}},
		{[]int{2, 4, 9, 3}, -2, []int{12, 5, 6, 13}},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%+v", test.code)
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.expected, decrypt(test.code, test.k))
		})
	}
}
