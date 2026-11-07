package threesumcloset

import "math"

func threeSumClosest(nums []int, target int) int {
	closetDiff := math.MaxInt32
	closetSum := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				var diff int
				if sum > target {
					diff = sum - target
				} else {
					diff = target - sum
				}

				if diff < closetDiff {
					closetDiff = diff
					closetSum = sum
				}
			}
		}
	}
	return closetSum
}
