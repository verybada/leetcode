package addbinary

import "fmt"

func addBinary(a string, b string) string {
	lengthA := len(a)
	lengthB := len(b)
	diff := lengthA - lengthB
	switch {
	case diff > 0:
		b = fmt.Sprintf("%0*s", lengthA, b)
	case diff < 0:
		a = fmt.Sprintf("%0*s", lengthB, a)
	}

	answer := make([]byte, len(a))
	carry := 0
	for i := len(a) - 1; i >= 0; i-- {
		binaryA := int(a[i] - '0')
		binaryB := int(b[i] - '0')
		sum := binaryA + binaryB + carry
		answer[i] = byte(sum%2) + '0'
		carry = sum / 2
	}

	if carry == 1 {
		answer = append([]byte{'1'}, answer...)
	}

	return string(answer)
}

func addBinary2(a string, b string) string {
	lengthA := len(a)
	lengthB := len(b)
	diff := lengthA - lengthB
	switch {
	case diff > 0:
		b = fmt.Sprintf("%0*s", lengthA, b)
	case diff < 0:
		a = fmt.Sprintf("%0*s", lengthB, a)
	}

	answer := make([]byte, len(a))
	carry := false
	for i := len(a) - 1; i >= 0; i-- {
		binaryA := a[i]
		binaryB := b[i]
		if binaryA != binaryB {
			answer[i] = '1'
			if carry {
				answer[i] = '0'
				carry = true
			}
		} else {
			answer[i] = '0'
			if carry {
				answer[i] = '1'
				carry = binaryA == '1'
			}

			if binaryA == '1' {
				carry = true
			}
		}
	}

	if carry {
		answer = append([]byte{'1'}, answer...)
	}

	return string(answer)
}
