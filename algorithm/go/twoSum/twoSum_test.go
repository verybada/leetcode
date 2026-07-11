package twoSum

import (
	"reflect"
	"testing"
)

var tests = []struct {
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
	{
		[]int{1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 7, 1, 1, 1, 1, 1},
		11,
		[]int{5, 11},
	},
}

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		output := twoSum(test.input, test.target)
		if !reflect.DeepEqual(output, test.output) {
			t.Fatal()
		}
	}
}

func TestTwoSum2(t *testing.T) {
	for _, test := range tests {
		output := twoSum2(test.input, test.target)
		if !reflect.DeepEqual(output, test.output) {
			t.Fatal()
		}
	}
}
