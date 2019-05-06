class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        self.head = None
        self.current = None
        def put(node):
            if not self.head:
                self.head = node
                self.current = self.head
            else:
                self.current.next = node
                self.current = node

        while True:
            if not l1:
                put(l2)
                break
            elif not l2:
                put(l1)
                break

            if l1.val > l2.val:
                put(l2)
                l2 = l2.next
            elif l1.val <= l2.val:
                put(l1)
                l1 = l1.next
        return self.head
