package queriesOnAPermutationWithKey

import ()

func processQueries(queries []int, m int) []int {
	result := make([]int, 0, len(queries))
	P := makeP(m)
	for _, value := range queries {
		pValue := getPosition(P, value)
		result = append(result, pValue)
	}
	return result
}

func makeP(m int) []int {
	P := make([]int, m)
	for i := 0; i < m; i++ {
		P[i] = i + 1
	}
	return P
}

func getPosition(P []int, value int) int {
	lastPValue := 0
	for index, pValue := range P {
		if index > 0 {
			P[index] = lastPValue
		}

		if pValue == value {
			P[0] = pValue
			return index
		}

		lastPValue = pValue
	}

	return -1
}
