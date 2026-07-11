package reversenodesinkgroup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		input []int
		k     int
		want  []int
	}{
		{
			[]int{},
			1,
			[]int{},
		},
		{
			[]int{1, 2, 3, 4, 5},
			1,
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4},
			2,
			[]int{2, 1, 4, 3},
		},
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{2, 1, 4, 3, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			3,
			[]int{3, 2, 1, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			5,
			[]int{5, 4, 3, 2, 1},
		},
		{
			[]int{1, 2, 3, 4, 5},
			6,
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, test := range tests {
		got := reverseKGroup(generateList(test.input), test.k)
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
