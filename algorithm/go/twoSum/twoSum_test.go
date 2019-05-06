package twoSum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		output []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{-3, 0, 3, 90},
			0,
			[]int{0, 2},
		},
	}

	for _, test := range tests {
		output := twoSum(test.input, test.target)
		if !reflect.DeepEqual(output, test.output) {
			t.Fatal()
		}
	}
}
