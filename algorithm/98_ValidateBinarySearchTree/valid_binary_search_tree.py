# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right

class Solution:
    def __init__(self):
        self.results = list()

    def _walk(self, node):
        if node.left:
            self._walk(node.left)

        self.results.append(node.val)

        if node.right:
            self._walk(node.right)

    def isValidBST(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if not root:
            return True

        self._walk(root)
        prev = None
        for val in self.results:
            if prev is None:
                prev = val
            elif prev >= val:
                return False
            prev = val
        return True


if __name__ == '__main__':
    node1 = TreeNode(1)
    node3 = TreeNode(3)
    root = TreeNode(2, node1, node3)
    assert Solution().isValidBST(root) is True

    node1 = TreeNode(1)
    node3 = TreeNode(3)
    node6 = TreeNode(6)
    node4 = TreeNode(4, node3, node6)
    root = TreeNode(5, node1, node4)
    assert Solution().isValidBST(root) is False
