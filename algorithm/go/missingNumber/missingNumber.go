package missingNumber

import (
    "sort"
)

func missingNumber(nums []int) int {
    numsLen := len(nums)
    newNums := make([]int, numsLen+1, numsLen+1)

    for _, value := range nums {
        newNums[value] = 1
    }

    for index, value := range newNums {
        if value == 0 {
            return index
        }
    }

    // should not be here
    return -1
}

/*
func missingNumber(nums []int) int {
    numsLen := len(nums)
    sort.Ints(nums)
    for index, value := range nums {
        if index != value {
            return index
        }
    }

    return numsLen
}
*/
