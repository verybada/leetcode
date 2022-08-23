from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def swapNodes(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        length = 0
        node = head
        while True:
            length += 1
            if not node.next:
                break
            node = node.next

        left_node = None
        right_node = None
        node = head
        for index in range(length):
            if index == k - 1:
                left_node = node

            elif index == length - k:
                right_node = node

            node = node.next

        if not left_node or not right_node:
            return head
        left_node.val, right_node.val = right_node.val, left_node.val
        return head
