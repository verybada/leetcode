package relativeSortArray

import (
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	weightMap := make(map[int]int)
	for index, value := range arr2 {
		weightMap[value] = index
	}

	sort.Slice(arr1, func(i int, j int) bool {
		left := arr1[i]
		right := arr1[j]

		leftWeight, leftOK := weightMap[left]
		rightWeight, rightOK := weightMap[right]

		if leftOK && rightOK {
			return leftWeight <= rightWeight
		} else if leftOK && !rightOK {
			return true
		} else if !leftOK && rightOK {
			return false
		} else {
			return left <= right
		}
	})
	return arr1
}
