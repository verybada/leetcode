# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        first = head
        max_len = 0
        node = head
    
        # find max length
        while node:
            max_len += 1
            node = node.next

        target_pos = max_len-n-1
        # if targetpos < 0, means remove head
        if target_pos < 0:
          return head.next
          
        # find our target
        node = head
        while target_pos > 0:
          node = node.next
          target_pos -= 1
          
        # remove it
        target = node.next
        if target:
          node.next = target.next
        else:
          node.next = None
        
        return head
