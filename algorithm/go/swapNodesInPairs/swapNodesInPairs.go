package swapNodesInPairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	top := head
	odd := true
	var previous *ListNode
	current := head
	for current != nil {
		if !odd {
			previous.Val, current.Val = current.Val, previous.Val
		}
		previous = current
		current = current.Next
		odd = !odd
	}
	return top
}
