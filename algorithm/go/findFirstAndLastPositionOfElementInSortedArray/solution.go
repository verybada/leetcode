package findFirstAndLastPositionOfElementInSortedArray

func searchRange(nums []int, target int) []int {
	return []int{
		findPosition(nums, target, true),
		findPosition(nums, target, false),
	}
}

func findPosition(nums []int, target int, isLeft bool) int {
	position := -1
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			position = mid
			if isLeft {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return position
}

func searchRangeOLogN(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := 0
	leftHit := false
	right := len(nums) - 1
	rightHit := false
	for left <= right {
		if !leftHit {
			if nums[left] == target {
				leftHit = true
			} else {
				left++
			}
		}

		if !rightHit {
			if nums[right] == target {
				rightHit = true
			} else {
				right--
			}
		}

		if leftHit && rightHit {
			break
		}
	}

	if leftHit && rightHit {
		return []int{left, right}
	}

	return []int{-1, -1}
}
