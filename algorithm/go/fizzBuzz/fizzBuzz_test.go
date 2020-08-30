package fizzBuzz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FizzBuzz(t *testing.T) {
	expected := []string{
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz",
	}

	require.Equal(t, expected, fizzBuzz(15))
}
