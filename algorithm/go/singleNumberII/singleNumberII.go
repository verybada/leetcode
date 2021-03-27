package singlenumberii

func singleNumberII(nums []int) int {
	m := make(map[int]int, 0)
	for _, num := range nums {
		count, ok := m[num]
		if !ok {
			m[num] = 1
		}
		m[num] = count + 1
	}

	for num, count := range m {
		if count == 1 {
			return num
		}
	}

	return -1
}
