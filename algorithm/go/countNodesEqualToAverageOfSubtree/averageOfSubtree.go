package averageOfSubtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	count := 0
	sumSubTree(root, &count)
	return count
}

func sumSubTree(node *TreeNode, count *int) (int, int) {
	if node == nil {
		return 0, 0
	}
	leftSum, leftCount := sumSubTree(node.Left, count)
	rightSum, rightCount := sumSubTree(node.Right, count)
	sum := (leftSum + rightSum + node.Val)
	totalCount := (leftCount + rightCount + 1)
	if sum/totalCount == node.Val {
		*count += 1
	}
	return sum, totalCount
}
