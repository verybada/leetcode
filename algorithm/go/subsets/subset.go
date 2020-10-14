package subset

func subsets(nums []int) [][]int {
	subsets := make([][]int, 0)
	subsets = append(subsets, []int{})
	for _, num := range nums {
		for _, subset := range subsets {
			newSubset := append([]int{}, subset...) // make new slice
			newSubset = append(newSubset, num)
			subsets = append(subsets, newSubset)
		}
	}

	return subsets
}
