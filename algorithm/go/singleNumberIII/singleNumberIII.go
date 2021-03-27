package singlenumberiii

func singleNumberIII(nums []int) []int {
	m := make(map[int]int, 0)
	for _, num := range nums {
		_, ok := m[num]
		if ok {
			delete(m, num)
		} else {
			m[num] = 1
		}
	}

	results := make([]int, 0, len(m))
	for num, _ := range m {
		results = append(results, num)
	}

	return results
}
