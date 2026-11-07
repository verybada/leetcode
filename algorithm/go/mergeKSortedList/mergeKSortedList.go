package mergeksortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	heap := []*ListNode{}
	for _, list := range lists {
		heap = append(heap, list)

		if len(heap) == 2 {
			newList := mergeTwoLists(heap[0], heap[1])
			heap = []*ListNode{newList}
		}
	}

	return mergeTwoLists(heap[0], nil)
}

func mergeTwoLists(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	head := &ListNode{}
	current := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
		current = current.Next
	}

	if left != nil {
		current.Next = left
	} else if right != nil {
		current.Next = right
	}

	return head.Next
}
