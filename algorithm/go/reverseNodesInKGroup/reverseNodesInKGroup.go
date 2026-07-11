package reversenodesinkgroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	top := head
	var previousGroupTail *ListNode
	stack := make([]*ListNode, 0, k)
	current := head
	for current != nil {
		if len(stack) == k {
			for i := len(stack) - 1; i >= 0; i-- {
				node := stack[i]
				if i == len(stack)-1 {
					// group head
					if previousGroupTail != nil {
						previousGroupTail.Next = node
					}
					node.Next = stack[i-1]
				} else if i != 0 {
					node.Next = stack[i-1]
				} else {
					// group tail
					node.Next = current
					if top == node {
						top = stack[len(stack)-1]
					}
					previousGroupTail = node
				}

			}
			stack = stack[:0]
		}

		stack = append(stack, current)
		current = current.Next
	}

	if len(stack) == k {
		for i := len(stack) - 1; i >= 0; i-- {
			node := stack[i]
			if i == len(stack)-1 {
				// group head
				if previousGroupTail != nil {
					previousGroupTail.Next = node
				}
				node.Next = stack[i-1]
			} else if i != 0 {
				node.Next = stack[i-1]
			} else {
				// group tail
				node.Next = current
				if top == node {
					top = stack[len(stack)-1]
				}
				previousGroupTail = node
			}

		}
		stack = stack[:0]
	}

	return top
}
