package elementAppearingMoreThanOneQuarterInSortedArray

func findSpecialInteger(arr []int) int {
	arrLen := len(arr)
	target := (arrLen / 4)

	currentValue := 0
	count := 0
	for _, value := range arr {
		if value != currentValue {
			currentValue = value
			count = 1
		} else {
			count += 1
		}

		if count > target {
			return value
		}
	}
	return -1
}
