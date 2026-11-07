package combinationsum

func combinationSum(candidates []int, target int) [][]int {
	answer := [][]int{}
	for i := 0; i < len(candidates); i++ {
		candidate := candidates[i]
		newTarget := target - candidate
		newCandidates := calculate(candidates[i:], newTarget, []int{candidate})
		if newCandidates == nil {
			continue
		}

		answer = append(answer, newCandidates...)
	}
	return answer
}

func calculate(candidates []int, target int, selected []int) [][]int {
	if target == 0 {
		return [][]int{selected}
	}

	if target < 0 {
		return nil
	}

	var answers [][]int
	for i, candidate := range candidates {
		newTarget := target - candidate
		newSelected := append([]int{}, selected...)
		newSelected = append(newSelected, candidate)
		answer := calculate(candidates[i:], newTarget, newSelected)
		if answer == nil {
			continue
		}
		answers = append(answers, answer...)
	}
	return answers
}
