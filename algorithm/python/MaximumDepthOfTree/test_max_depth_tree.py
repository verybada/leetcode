from .max_depth_tree import Solution, Node


def test_max_depth_none():
    result = Solution().maxDepth(None)
    assert result == 0


def test_max_depth_root():
    root = Node(1)
    result = Solution().maxDepth(root)
    assert result == 1


def test_max_depth_example1():
    node2 = Node(2)
    node4 = Node(4)
    node5 = Node(5)
    node6 = Node(6)
    node3 = Node(3, children=[node5, node6])
    root = Node(1, children=[node2, node4, node3])
    result = Solution().maxDepth(root)
    assert result == 3


def test_max_depth_example2():
    node14 = Node(14)
    node13 = Node(13)
    node12 = Node(12)
    node11 = Node(11, children=[node14, ])
    node10 = Node(10)
    node9 = Node(9, children=[node13, ])
    node8 = Node(8, children=[node12, ])
    node7 = Node(7, children=[node11, ])
    node6 = Node(6)
    node5 = Node(5, children=[node9, node10])
    node4 = Node(4, children=[node8, ])
    node3 = Node(3, children=[node6, node7])
    node2 = Node(2)
    root = Node(1, children=[node2, node3, node4, node5])
    result = Solution().maxDepth(root)
    assert result == 5
