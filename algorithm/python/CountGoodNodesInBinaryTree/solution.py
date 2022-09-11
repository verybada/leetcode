from __future__ import annotations


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def goodNodes(self, root: TreeNode, max_value=-10000) -> int:
        max_value = max(root.val, max_value)
        count = 1 if root.val >= max_value else 0
        if root.left:
            count += self.goodNodes(root.left, max_value)
        if root.right:
            count += self.goodNodes(root.right, max_value)
        return count
