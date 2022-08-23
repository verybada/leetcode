class Solution:
    def numSplits(self, s: str) -> int:
        good_split = 0
        left = dict()
        right = dict()
        for c in s:
            count = right.get(c, 0)
            count += 1
            right[c] = count

        for c in s:
            count = left.get(c, 0)
            left[c] = count + 1

            count = right[c]
            count -= 1
            if count != 0:
                right[c] = count
            else:
                del right[c]

            if len(left) == len(right):
                good_split += 1

        return good_split
