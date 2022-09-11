from .solution import TreeNode, Solution


def test_good_nodes():
    left_one_three = TreeNode(3)
    left_one = TreeNode(1, left=left_one_three)
    right_four_one = TreeNode(1)
    right_four_five = TreeNode(5)
    right_four = TreeNode(4, left=right_four_one, right=right_four_five)
    root = TreeNode(3, left=left_one, right=right_four)
    expected = 4
    assert Solution().goodNodes(root) == expected


def test_good_nodes_case2():
    left_three_four = TreeNode(4)
    left_three_two = TreeNode(2)
    left_three = TreeNode(3, left=left_three_four, right=left_three_two)
    root = TreeNode(3, left=left_three)
    expected = 3
    assert Solution().goodNodes(root) == expected
