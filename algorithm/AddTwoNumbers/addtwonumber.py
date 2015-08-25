# Comment ListNode when submit to leetcode
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def _sum(self, x, y):
        if x is None:
            return y
        elif y is None:
            return x
        return x+y

    def _add(self, node1, node2, carry):
        """
        :param node1: A node or None
        :type node1: ListNode or None
        :param node2: A node or None
        :type node2: ListNode or None
        :param int carry: a positive int (include 0) carry
        :return A node or None
        :rtype ListNode or None
        """
        if node1 is None and node2 is None:
            if carry != 0:
                return ListNode(carry)
            else:
                return None

        # nodes might be None
        v1 = node1.val if node1 else None
        n1 = node1.next if node1 else None
        v2 = node2.val if node2 else None
        n2 = node2.next if node2 else None

        # create current node
        node = ListNode(self._sum(v1, v2))
        # add carry to current value
        node.val += carry
        # create next node from next value and current carry
        node.next = self._add(n1, n2, node.val/10)
        # make sure current value stay in correct range
        node.val %= 10
        return node

    def addTwoNumbers(self, l1, l2):
        """
        :param ListNode l1
        :param ListNode l2
        :return A ListNode which is sum(l1, l2)
        :rtype: ListNode
        >>> _ = Run((2,4,3), (5,6,4))
        [7, 0, 8]

        >>> _ = Run((5,), (5,))
        [0, 1]

        >>> _ = Run((1,8), (0,))
        [1, 8]
        """
        return self._add(l1, l2, 0)
