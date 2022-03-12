package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustomSortString(t *testing.T) {
	tests := []struct {
		order    string
		s        string
		expected string
	}{
		{"cba", "abcd", "cbad"},
		{"cbafg", "abcd", "cbad"},
		{"abc", "def", "def"},
		{"abc", "cba", "abc"},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			result := customSortString(test.order, test.s)
			require.Equal(t, test.expected, result)
		})
	}
}
