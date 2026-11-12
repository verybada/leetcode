package plusone

func plusOne(digits []int) []int {
	length := len(digits)
	over := 0
	for i := length - 1; i >= 0; i-- {
		number := digits[i]
		if i == length-1 {
			number += 1
		}

		if over != 0 {
			number += over
			over = 0
		}

		if number <= 9 {
			digits[i] = number
			return digits
		}

		digits[i] = 0
		over = 1
	}

	return append([]int{1}, digits...)
}
