package fizzBuzz

import (
	"fmt"
)

func fizzBuzz(n int) []string {
	output := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			output = append(output, "FizzBuzz")
		} else if i%3 == 0 {
			output = append(output, "Fizz")
		} else if i%5 == 0 {
			output = append(output, "Buzz")
		} else {
			output = append(output, fmt.Sprintf("%d", i))
		}
	}
	return output
}
