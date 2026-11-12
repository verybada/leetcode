package removeduplicates

func removeDuplicates2(nums []int) int {
	currentNum := nums[0]
	lastBoundary := 0
	count := 1
	for i := 1; i < len(nums); i++ {
		num := nums[i]

		if currentNum != num {
			count = 1
			currentNum = num

			nums = append(nums[:lastBoundary+1], nums[i:]...)
			i = lastBoundary + 1
		} else {
			count += 1
		}

		if count > 2 {
			continue
		}

		lastBoundary = i
	}
	return lastBoundary + 1
}
