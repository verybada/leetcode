package singleNumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleNumber(t *testing.T) {
	methods := []struct {
		name string
		f    func([]int) int
	}{
		{
			"map",
			singleNumberMap,
		},
		{
			"xor",
			singleNumberXor,
		},
		{
			"sort",
			singleNumberSort,
		},
	}
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{1},
			1,
		},
		{
			[]int{2, 2, 1},
			1,
		},
		{
			[]int{4, 1, 2, 1, 2},
			4,
		},
	}

	for _, method := range methods {
		t.Run(method.name, func(t *testing.T) {
			for _, test := range tests {
				name := fmt.Sprintf("%+v", test.input)
				t.Run(name, func(t *testing.T) {
					output := method.f(test.input)
					require.Equal(t, test.expected, output)
				})
			}
		})
	}
}
