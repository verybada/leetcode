package twoSum

func twoSum(nums []int, target int) []int {
	cache := map[int]int{}
	for index, num := range nums {

		if cachedIndex, ok := cache[num]; ok {
			r := make([]int, 0, 2)
			r = append(r, cachedIndex)
			r = append(r, index)
			return r
		}

		diff := target - num
		cache[diff] = index
	}

	return []int{}
}
