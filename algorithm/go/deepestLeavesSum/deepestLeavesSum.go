package deepestLeavesSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	maxDepth := 1
	sum := root.Val
	findLeave(root.Left, &sum, &maxDepth, 2)
	findLeave(root.Right, &sum, &maxDepth, 2)
	return sum
}

func findLeave(node *TreeNode, sum *int, maxDepth *int, depth int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		if depth == *maxDepth {
			*sum += node.Val
		} else if depth > *maxDepth {
			*maxDepth = depth
			*sum = node.Val
		}
		return
	}

	findLeave(node.Left, sum, maxDepth, depth+1)
	findLeave(node.Right, sum, maxDepth, depth+1)
}
