package reduceArraySizeToTheHalf

import (
	"sort"
)

func minSetSize(arr []int) int {
	length := len(arr)
	valueCountMap := make(map[int]int)
	for _, value := range arr {
		count, ok := valueCountMap[value]
		if !ok {
			count = 0
		}
		valueCountMap[value] = count + 1
	}

	counts := make([]int, 0)
	for _, count := range valueCountMap {
		counts = append(counts, count)
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] >= counts[j]
	})

	half := length / 2
	for index, count := range counts {
		length -= count
		if length <= half {
			return index + 1
		}
	}

	return -1
}
