package getAllelements

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllElements(t *testing.T) {
	// [1, 2, 4]
	node_1 := TreeNode{Val: 1}
	node_4 := TreeNode{Val: 4}
	root1 := TreeNode{Val: 2, Left: &node_1, Right: &node_4}

	// [0, 1, 3]
	node_0 := TreeNode{Val: 0}
	node_3 := TreeNode{Val: 3}
	root2 := TreeNode{Val: 1, Left: &node_0, Right: &node_3}

	expected := []int{0, 1, 1, 2, 3, 4}
	require.Equal(t, expected, getAllElements(&root1, &root2))
}
