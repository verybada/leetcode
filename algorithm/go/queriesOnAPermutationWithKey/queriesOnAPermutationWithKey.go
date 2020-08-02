package queriesOnAPermutationWithKey

import ()

func processQueries(queries []int, m int) []int {
	result := make([]int, 0, len(queries))
	P := makeP(m)
	for index, value := range queries {
		pValue := getPosition(P, value)
		result = append(result, pValue)
	}
	return result
}

func makeP(m int) []int {
	P := make([]int, 0, m)
	for i := 0; i < m; i++ {
		P[i] = i
	}
	return P
}

func getPosition(P []int, value int) int {
	for index, pValue := range P {
		if index > 0 {
			P[index] = P[index-1]
		}

		if pValue == value {
			P[0] = pValue
			return index
		}
	}
}
