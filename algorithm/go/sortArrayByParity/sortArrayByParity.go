package sortArrayByParity

import (
)

func sortArrayByParity(A []int) []int {
    oddIndex := -1
    for index, value := range A {
        if value % 2 == 0 {
            if oddIndex >= 0 {
                A[index], A[oddIndex] = A[oddIndex], A[index]
                oddIndex += 1
            }
        } else {
            if oddIndex == -1 {
                oddIndex = index
            }
        }
    }
    return A
}
