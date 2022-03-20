from .max_depth import Solution, TreeNode


def test_max_depth_none():
    assert 0 == Solution().maxDepth(None)


def test_max_depth_root():
    root = TreeNode(1)
    assert 1 == Solution().maxDepth(root)


def test_max_depth_example1():
    node7 = TreeNode(7)
    node15 = TreeNode(15)
    node20 = TreeNode(20, left=node15, right=node7)
    node9 = TreeNode(9)
    root = TreeNode(3, left=node9, right=node20)
    assert 3 == Solution().maxDepth(root)


def test_max_depth_example2():
    node2 = TreeNode(2)
    root = TreeNode(1, right=node2)
    assert 2 == Solution().maxDepth(root)
