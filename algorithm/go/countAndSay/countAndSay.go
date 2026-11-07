package countAndSay

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := countAndSay(n - 1)
	return toReadString(prev)
}

func toReadString(input string) string {
	output := ""
	buf := []rune{}
	for _, char := range input {
		if len(buf) == 0 {
			buf = append(buf, char)
			continue
		}

		if char == buf[0] {
			buf = append(buf, char)
			continue
		}

		output += fmt.Sprintf("%d%s", len(buf), string(buf[0]))
		buf = []rune{char}
	}
	output += fmt.Sprintf("%d%s", len(buf), string(buf[0]))
	return output
}
