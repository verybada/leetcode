package sortArrayByParity

import (
    "testing"
    "reflect"
)

func TestSortArrayByParity(t *testing.T) {
    tests := []struct {
        input []int
        output []int
    } {
        {
            []int {0},
            []int {0},
        },
        {
            []int {0, 2},
            []int {0, 2},
        },
        {
            []int {3, 1, 2, 4},
            []int {2, 4, 3, 1},
        },
        {
            []int {1, 2, 3, 4, 5, 6, 7, 8, 9},
            []int {2, 4, 6, 8, 5, 3, 7, 1, 9},
        },
    }

    for _, test := range tests {
        expect := sortArrayByParity(test.input)
        if !reflect.DeepEqual(test.output, expect) {
            t.Fatalf("input: %v, expect: %v, result: %v",
                test.input, test.output, expect)
        }
    }
}
