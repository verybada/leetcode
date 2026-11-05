package romantoint

import "fmt"

func romanToInt(input string) int {
	mapper := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	value := 0
	i := 0
	for i < len(input) {
		left := mapper[string(input[i])]
		if i == len(input)-1 {
			value += left
			break
		}

		right := mapper[string(input[i+1])]
		fmt.Printf("left: %d, right: %d\n", left, right)
		if left >= right {
			value += left
			i += 1
		} else {
			value += (right - left)
			i += 2
		}

	}
	return value
}
