class Solution:
    def numJewelsInStones(self, J, S):
        """
        :type J: str
        :type S: str
        :rtype: int
        """
        maps = dict()
        for c in J:
            maps[c] = c
        count = 0
        for c in S:
            if c in maps:
                count += 1
        return count
