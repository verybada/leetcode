package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinOperations(t *testing.T) {
	tests := []struct {
		logs     []string
		expected int
	}{
		{[]string{"d1/", "d2/", "../", "d21/", "./"}, 2},
		{[]string{"d1/", "../", "../", "../"}, 0},
		{[]string{"../", "../"}, 0},
		{[]string{"a/", "b/", "c/"}, 3},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v", test.logs)
		t.Run(name, func(t *testing.T) {
			result := minOperations(test.logs)
			require.Equal(t, test.expected, result)
		})
	}
}
