package main

func customSortString(order string, s string) string {
	orderCount := make(map[rune]int)
	for _, char := range order {
		orderCount[char] = 0
	}

	others := make([]rune, 0)
	for _, char := range s {
		count, ok := orderCount[char]
		if !ok {
			others = append(others, char)
			continue
		}

		count++
		orderCount[char] = count
	}

	sortedString := ""
	for _, char := range order {
		count := orderCount[char]
		for i := 0; i < count; i++ {
			sortedString += string(char)
		}
	}
	sortedString += string(others)
	return sortedString
}
