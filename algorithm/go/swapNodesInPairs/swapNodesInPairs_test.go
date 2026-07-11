package swapNodesInPairs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{2, 1, 4, 3},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{2, 1, 4, 3, 5},
		},
	}
	for _, test := range tests {
		got := swapPairs(generateList(test.input))
		assertList(t, got, test.want)
	}
}

func generateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

func assertList(t *testing.T, got *ListNode, want []int) {
	gotList := []int{}
	current := got
	for current != nil {
		gotList = append(gotList, current.Val)
		current = current.Next
	}
	require.Equal(t, want, gotList)
}
