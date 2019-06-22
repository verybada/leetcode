package flippingAnImage

import (
    "reflect"
    "testing"
)

func TestFlipAndInvertImage(t *testing.T) {
    tests := []struct{
        input [][]int
        output [][]int
    } {
        {
            input: [][]int{
                []int {1, 1, 0},
                []int {1, 0, 1},
                []int {0, 0, 0},
            },
            output: [][]int{
                []int {1, 0, 0},
                []int {0, 1, 0},
                []int {1, 1, 1},
            },
        },
        {
            input: [][]int{
                []int {1, 1, 0, 0},
                []int {1, 0, 0, 1},
                []int {0, 1, 1, 1},
                []int {1, 0, 1, 0},
            },
            output: [][]int{
                []int {1, 1, 0, 0},
                []int {0, 1, 1, 0},
                []int {0, 0, 0, 1},
                []int {1, 0, 1, 0},
            },
        },
    }

    for _ ,test := range tests {
        result := flipAndInvertImage(test.input)
        if !reflect.DeepEqual(result, test.output) {
            t.Fatalf("input: %v, expect: %v, result: %v",
                test.input, test.output, result)
        }
    }
}

