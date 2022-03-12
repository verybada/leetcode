package main

import (
	"sort"
)

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	count := 0
	for _, cost := range costs {
		coins -= cost
		if coins < 0 {
			return count
		}
		count++
	}
	return count
}
