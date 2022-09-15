package averageOfSubtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAverageOfSubtree(t *testing.T) {
	node6 := &TreeNode{Val: 6}
	node5 := &TreeNode{Val: 5, Right: node6}
	node0 := &TreeNode{Val: 0}
	node1 := &TreeNode{Val: 1}
	node8 := &TreeNode{Val: 8, Left: node0, Right: node1}
	root := &TreeNode{Val: 4, Left: node8, Right: node5}

	result := averageOfSubtree(root)
	require.Equal(t, 5, result)
}
