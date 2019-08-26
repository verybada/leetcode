package main

import (
	"testing"
)

func TestJudgeCircle(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"LL", false},
		{"UD", true},
	}

	for _, test := range tests {
		result := judgeCircle(test.input)
		if result != test.output {
			t.Fatalf("%s expected %v but %v", test.input, test.output, result)
		}
	}
}
