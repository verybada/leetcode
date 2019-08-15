package main

import (
	"testing"
)

func TestRepeatedNTimes(t *testing.T) {
	tests := []struct {
		A     []int
		value int
	}{
		{
			[]int{1, 2, 3, 3},
			3,
		},
		{
			[]int{7, 6, 6, 8},
			6,
		},
		{
			[]int{1, 2, 2, 3},
			2,
		},
	}

	for _, test := range tests {
		if value := repeatedNTimes(test.A); value != test.value {
			t.Fatalf("array %v, expected value %d, but got %d",
				test.A, test.value, value)
		}
	}
}
