package missingNumber

import (
    "testing"
)

func TestMissingNumber(t *testing.T) {
    nums := []int {3, 0, 1}
    if missingNumber(nums) != 2 {
        t.Fatalf("Not equal 2")
    }

    nums = []int {0}
    if missingNumber(nums) != 1 {
        t.Fatalf("Not equal 1")
    }

    nums = []int {1}
    if missingNumber(nums) != 0 {
        t.Fatalf("Not equal 0")
    }
}
