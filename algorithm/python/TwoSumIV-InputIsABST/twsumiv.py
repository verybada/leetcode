class Solution:
    def _cal(self, node, k, maps):
        target = k - node.val
        if target in maps:
            return True
        maps[node.val] = node.val
        
        if node.left and self._cal(node.left, k, maps):
            return True                
        if node.right and self._cal(node.right, k, maps):
            return True
        return False
            
    def findTarget(self, root, k):
        """
        :type root: TreeNode
        :type k: int
        :rtype: bool
        """
        maps = dict()
        return self._cal(root, k, maps)
        
