package rangeSumOfBST

import ()

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	} else if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	} else if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	} else {
		return rangeSumBST(root.Left, L, R) + root.Val + rangeSumBST(root.Right, L, R)
	}
}
