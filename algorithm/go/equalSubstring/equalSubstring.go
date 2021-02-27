package equalsubstring

func equalSubstring(s string, t string, maxCost int) int {
	source := []rune(s)
	target := []rune(t)
	maxLength := 0
	left := 0
	cost := 0
	for right := 0; right < len(source); right++ {
		cost += abs(int(source[right] - target[right]))

		if cost > maxCost {
			cost -= abs(int(source[left] - target[left]))
			left++
		}
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

func abs(value int) int {
	if value >= 0 {
		return value
	}
	return -value
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
