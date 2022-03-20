from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0

        left_depth = right_depth = 1
        if root.left:
            left_depth = 1+self.maxDepth(root.left)

        if root.right:
            right_depth = 1+self.maxDepth(root.right)
        return max(left_depth, right_depth)
