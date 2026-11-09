package permuteii

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	answers := [][]int{}
	path := []int{}
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

			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
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
