package singleNumber

import (
	"sort"
)

func singleNumberXor(nums []int) int {
	x := 0
	for _, num := range nums {
		x ^= num
	}
	return x
}

func singleNumberMap(nums []int) int {
	tmp := make(map[int]bool, 0)
	for _, num := range nums {
		_, ok := tmp[num]
		if ok {
			delete(tmp, num)
		} else {
			tmp[num] = true
		}
	}

	for num, _ := range tmp {
		return num
	}

	return -1
}

func singleNumberSort(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	sort.Ints(nums)
	for i := 0; i < length; i += 2 {
		if i == 0 {
			if nums[i] != nums[i+1] {
				return nums[i]
			}
		} else if i == length-1 {
			if nums[i-1] != nums[i] {
				return nums[i]
			}
		} else {
			if nums[i] != nums[i+1] && nums[i-1] != nums[i] {
				return nums[i]
			}
		}
	}
	return -1
}
