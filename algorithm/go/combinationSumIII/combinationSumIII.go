package combinationSumIII

func combinationSum3(k int, n int) [][]int {
	results := [][]int{}
	helper(k, n, 1, []int{}, &results)
	return results
}

func helper(
	remaining int, target int, currentValue int,
	result []int, results *[][]int,
) {
	if remaining == 0 {
		if target == 0 {
			*results = append(*results, result)
		}
		return
	}
	if currentValue > target {
		return
	}

	if currentValue > 9 {
		return
	}

	nextValue := currentValue + 1
	newResult := append([]int{}, result...)
	newResult = append(newResult, currentValue)
	helper(remaining-1, target-currentValue, nextValue, newResult, results)
	helper(remaining, target, nextValue, result, results)
}
