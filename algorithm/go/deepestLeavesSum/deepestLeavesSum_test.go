package deepestLeavesSum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeepestLeavesSum(t *testing.T) {
	node8 := &TreeNode{8, nil, nil}
	node6 := &TreeNode{6, nil, node8}
	node3 := &TreeNode{3, nil, node6}
	node7 := &TreeNode{7, nil, nil}
	node4 := &TreeNode{4, node7, nil}
	node5 := &TreeNode{5, nil, nil}
	node2 := &TreeNode{2, node4, node5}
	root := &TreeNode{1, node2, node3}

	result := deepestLeavesSum(root)
	require.Equal(t, 15, result)
}
