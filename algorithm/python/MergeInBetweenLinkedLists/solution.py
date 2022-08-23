# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def __str__(self) -> str:
        return f"{self.val}"


class Solution:
    def mergeInBetween(self, list1: ListNode, a: int, b: int, list2: ListNode) -> ListNode:
        index = 0
        list2_last_node = self._get_last_node(list2)
        head = list1
        left_bound = None
        node = list1
        while True:
            if index == a - 1:
                left_bound = node
                pass

            if index == b + 1:
                list2_last_node.next = node

            if not node.next:
                break
            node = node.next
            index += 1

        left_bound.next = list2
        return head

    def _get_last_node(self, node: ListNode) -> ListNode:
        while True:
            if not node.next:
                return node
            node = node.next
