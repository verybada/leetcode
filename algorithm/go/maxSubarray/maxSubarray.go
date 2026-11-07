package maxsubarray

func maxSubArray(nums []int) int {
	maxSum := 0
	currentSum := 0
	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]

		if currentSum > maxSum || i == 0 {
			maxSum = currentSum
		}

		if currentSum < 0 {
			currentSum = 0
		}
	}
	return maxSum
}

func maxSubArrayBigOMxN(nums []int) int {
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum > maxSum || i == 0 {
			maxSum = sum
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}
