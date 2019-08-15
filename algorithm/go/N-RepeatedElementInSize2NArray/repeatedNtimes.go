package main

func repeatedNTimes(A []int) int {
	repeatedValue := map[int]int{}
	for _, value := range A {
		if _, ok := repeatedValue[value]; ok {
			return value
		}
		repeatedValue[value] = 1
	}

	return -1
}
