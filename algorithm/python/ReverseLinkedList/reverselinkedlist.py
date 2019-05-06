# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseList(self, node):
        """
        :type head: ListNode
        :rtype: ListNode
        """       
        prev_node = None
        while node:
            next_ = node.next
            node.next = prev_node
            prev_node = node
            node = next_
            
        return prev_node
        
