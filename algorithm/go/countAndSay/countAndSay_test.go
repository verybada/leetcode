package countAndSay

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
		{6, "312211"},
		{7, "13112221"},
		{8, "1113213211"},
		{9, "31131211131221"},
		{10, "13211311123113112211"},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			require.Equal(t, tt.want, countAndSay(tt.n))
		})
	}
}

func TestToReadString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "single digit", input: "1", want: "11"},
		{name: "one repeated group", input: "1111", want: "41"},
		{name: "two distinct digits", input: "21", want: "1211"},
		{name: "final single digit", input: "1112", want: "3112"},
		{name: "two repeated groups", input: "1122", want: "2122"},
		{name: "alternating digits", input: "1212", want: "11121112"},
		{name: "separated matching digits", input: "11211", want: "211221"},
		{name: "three groups", input: "111221", want: "312211"},
		{name: "zero digit", input: "00011", want: "3021"},
		{name: "multi digit count", input: "111111111111", want: "121"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, toReadString(tt.input))
		})
	}
}
