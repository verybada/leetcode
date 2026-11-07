package combinationsumii

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	answers := [][]int{}
	for i := 0; i < len(candidates); i++ {
		candidate := candidates[i]
		if i > 0 && candidate == candidates[i-1] {
			continue
		}

		newTarget := target - candidate
		subanswers := calculate(candidates[i+1:], newTarget, []int{candidate})
		if len(subanswers) == 0 {
			continue
		}

		answers = append(answers, subanswers...)
	}
	return answers
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
		if i > 0 && candidate == candidates[i-1] {
			continue
		}

		newSelected := append([]int{}, selected...)
		newSelected = append(newSelected, candidate)
		newTarget := target - candidate
		subanswers := calculate(candidates[i+1:], newTarget, newSelected)
		if subanswers == nil {
			continue
		}

		answers = append(answers, subanswers...)
	}
	return answers
}
