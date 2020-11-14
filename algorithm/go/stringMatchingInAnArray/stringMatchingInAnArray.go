package stringMatchingInAnArray

import (
	"strings"
)

func stringMatching(words []string) []string {
	substrings := make(map[string]int, 0)
	for index, word := range words {
		for anotherIndex, anotherWord := range words {
			if index == anotherIndex {
				continue
			}

			if len(word) < len(anotherWord) {
				continue
			}

			if strings.Contains(word, anotherWord) {
				substrings[anotherWord] = 1
			}
		}
	}

	results := make([]string, 0, len(substrings))
	for substring, _ := range substrings {
		results = append(results, substring)
	}
	return results
}
