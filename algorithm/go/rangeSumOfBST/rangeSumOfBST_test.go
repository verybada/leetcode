package rangeSumOfBST

import (
    "testing"
)

func TestRangeSumBST1(t *testing.T) {
    node3 := &TreeNode{3, nil, nil}
    node7 := &TreeNode{7, nil, nil}
    node5 := &TreeNode{5, node3, node7}
    node18 := &TreeNode{18, nil, nil}
    node15 := &TreeNode{15, nil, node18}
    root := &TreeNode{10, node5, node15}
    if 32 != rangeSumBST(root, 7, 15) {
        t.Fatalf("sum not 32")
    }
}
