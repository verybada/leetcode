package findLuckyIntegerinAnArray

func findLucky(arr []int) int {
	valueMap := make(map[int]int)
	for _, value := range arr {
		count, ok := valueMap[value]
		if !ok {
			count = 0
		}

		valueMap[value] = count + 1
	}

	biggestLuckyInt := -1
	for value, count := range valueMap {
		if value != count {
			continue
		}

		if value <= biggestLuckyInt {
			continue
		}

		biggestLuckyInt = value
	}

	return biggestLuckyInt
}
