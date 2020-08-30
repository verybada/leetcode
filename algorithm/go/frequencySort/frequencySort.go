package frequencySort

import (
	"sort"
	"strings"
)

func frequencySort(input string) string {
	pool := make(map[rune]int)
	for _, char := range input {
		count, ok := pool[char]
		if !ok {
			count = 0
		}

		count++
		pool[char] = count
	}

	output := make([]string, 0, len(pool))
	for char, count := range pool {
		chars := strings.Repeat(string(char), count)
		output = append(output, chars)
	}

	sort.Slice(output, func(i, j int) bool {
		return len(output[i]) > len(output[j])
	})
	return strings.Join(output, "")
}
