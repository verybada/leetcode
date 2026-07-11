package twoSum

func twoSum2(nums []int, target int) []int {
	mapper := map[int]int{}
	for index, num := range nums {
		anotherIndex, ok := mapper[num]
		if ok {
			return []int{anotherIndex, index}
		}

		remain := target - num
		mapper[remain] = index
	}

	return nil
}
