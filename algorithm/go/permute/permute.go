package permute

func permute(nums []int) [][]int {
	var answers [][]int
	var path []int
	used := make([]bool, len(nums))
	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			answer := make([]int, len(path))
			copy(answer, path)
			answers = append(answers, answer)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])

			backtrack()

			path = path[:len(path)-1]
			used[i] = false
		}

	}
	backtrack()
	return answers
}

func permuteExt(nums []int) [][]int {
	used := make([]bool, len(nums))
	return calculate(nums, []int{}, used)
}

func calculate(nums []int, selected []int, used []bool) [][]int {
	if len(selected) == len(nums) {
		return [][]int{selected}
	}

	answers := [][]int{}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		nextSelected := make([]int, len(selected))
		copy(nextSelected, selected)
		nextUsed := make([]bool, len(used))
		copy(nextUsed, used)
		nextSelected = append(nextSelected, nums[i])
		nextUsed[i] = true
		subanswers := calculate(nums, nextSelected, nextUsed)
		answers = append(answers, subanswers...)
	}
	return answers
}
