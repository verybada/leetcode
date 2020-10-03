package findAndReplacePattern

func findAndReplacePattern(words []string, pattern string) []string {
	results := make([]string, 0)
	int_pattern := convertPattern(pattern)
	for _, word := range words {
		if !compare(word, int_pattern) {
			continue
		}

		results = append(results, word)
	}
	return results
}

func convertPattern(str string) []int {
	pattern := make([]int, 0, len(str))
	count := 0
	cache := make(map[rune]int)
	for _, char := range str {
		value, ok := cache[char]
		if ok {
			pattern = append(pattern, value)
			continue
		}

		cache[char] = count
		pattern = append(pattern, count)
		count++
	}

	return pattern
}

func compare(str string, pattern []int) bool {
	if len(str) != len(pattern) {
		return false
	}

	count := 0
	cache := make(map[rune]int)
	for index, char := range str {
		value, ok := cache[char]
		if !ok {
			cache[char] = count
			value = count
			count++
		}

		if pattern[index] != value {
			return false
		}
	}

	return true
}
