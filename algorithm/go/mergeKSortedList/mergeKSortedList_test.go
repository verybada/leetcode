package mergeksortedlist

import "testing"

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{name: "nil input"},
		{name: "empty input", lists: [][]int{}},
		{name: "single empty list", lists: [][]int{nil}},
		{name: "all empty lists", lists: [][]int{nil, {}, nil}},
		{name: "single list", lists: [][]int{{1, 2, 3}}, want: []int{1, 2, 3}},
		{
			name:  "example",
			lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			want:  []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:  "two interleaved lists",
			lists: [][]int{{1, 3, 5}, {2, 4, 6}},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "empty lists mixed in",
			lists: [][]int{nil, {1, 3}, {}, {2}, nil},
			want:  []int{1, 2, 3},
		},
		{
			name:  "negative values and duplicates",
			lists: [][]int{{-5, -1, 0}, {-3, -1, 2}, {-1}},
			want:  []int{-5, -3, -1, -1, -1, 0, 2},
		},
		{
			name:  "unequal lengths",
			lists: [][]int{{1}, {2, 3, 4, 5}, {6}, {7, 8}},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Fatalf("mergeKLists panicked: %v", err)
				}
			}()
			lists := make([]*ListNode, len(tt.lists))
			for i, values := range tt.lists {
				lists[i] = generateList(values)
			}
			assertList(t, mergeKLists(lists), tt.want)
		})
	}
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name        string
		left, right []int
		want        []int
	}{
		{name: "both empty"},
		{name: "left empty", right: []int{1, 2}, want: []int{1, 2}},
		{name: "right empty", left: []int{1, 2}, want: []int{1, 2}},
		{
			name:  "interleaved",
			left:  []int{1, 3, 5},
			right: []int{2, 4, 6},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{name: "left remainder", left: []int{2, 3, 4}, right: []int{1}, want: []int{1, 2, 3, 4}},
		{name: "right remainder", left: []int{1}, right: []int{2, 3, 4}, want: []int{1, 2, 3, 4}},
		{
			name:  "negative values and duplicates",
			left:  []int{-3, -1, 2},
			right: []int{-2, -1, 2},
			want:  []int{-3, -2, -1, -1, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertList(t, mergeTwoLists(generateList(tt.left), generateList(tt.right)), tt.want)
		})
	}
}

func generateList(values []int) *ListNode {
	head := &ListNode{}
	tail := head
	for _, value := range values {
		tail.Next = &ListNode{Val: value}
		tail = tail.Next
	}
	return head.Next
}

func assertList(t *testing.T, got *ListNode, want []int) {
	t.Helper()
	for i, value := range want {
		if got == nil {
			t.Fatalf("list ended at index %d; want %d nodes", i, len(want))
		}
		if got.Val != value {
			t.Fatalf("node %d = %d; want %d", i, got.Val, value)
		}
		got = got.Next
	}
	// A bounded walk also catches extra nodes and cycles without hanging.
	if got != nil {
		t.Fatalf("unexpected node after %d values: %d", len(want), got.Val)
	}
}
