# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

    def __str__(self):
        left = self.left.val if self.left else None
        right = self.right.val if self.right else None
        return 'left %s val %s right %s' % (left, self.val, right)

class Solution:
    def _walk_tree(self, node):
        if node.left:
            self._walk_tree(node.left)
        if node.right:
            self._walk_tree(node.right)

        node.left, node.right = node.right, node.left

    def invertTree(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        if root:
            self._walk_tree(root)
        return root
