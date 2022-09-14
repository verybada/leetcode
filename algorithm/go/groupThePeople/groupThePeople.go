package groupThePeople

func groupThePeople(groupSizes []int) [][]int {
	result := [][]int{}
	groups := map[int][]int{}
	for index, groupSize := range groupSizes {
		group, ok := groups[groupSize]
		if !ok {
			group = make([]int, 0, groupSize)
		}

		group = append(group, index)
		if len(group) == groupSize {
			result = append(result, group)
			delete(groups, groupSize)
		} else {
			groups[groupSize] = group
		}
	}
	return result
}
