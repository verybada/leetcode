from .solution import ListNode, Solution


def test_merge_in_between():
    l4 = ListNode(4)
    l3 = ListNode(3, l4)
    l2 = ListNode(2, l3)
    l1 = ListNode(1, l2)
    l0 = ListNode(0, l1)

    r1 = ListNode(101)
    r0 = ListNode(100, r1)

    node = Solution().mergeInBetween(l0, 1, 4, r0)
    assert _to_str(node) == "0->100->101"


def _to_str(node: ListNode) -> str:
    parts = list()
    while True:
        parts.append(f"{node.val}")
        if not node.next:
            break
        node = node.next
    return "->".join(parts)
