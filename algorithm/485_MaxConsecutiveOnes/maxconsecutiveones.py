class Solution:
    def findMaxConsecutiveOnes(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        max_count = 0
        count = 0
        for n in nums:
            if n == 0:
                max_count = max(count, max_count)
                count = 0
                continue
                
            count += 1
        max_count = max(count, max_count)
        return max_count
