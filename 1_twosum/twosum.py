class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]

        >>> obj = Solution()
        >>> obj.twoSum([2, 5, 7, 11], 9)
        [1, 3]
        >>> obj.twoSum([-1, 0, 1, 99], 0)
        [1, 3]
        >>> obj.twoSum([0, 1, 3, 0], 0)
        [1, 4]
        """

        hash = dict()
        for index in xrange(0, len(nums)):
            n = nums[index]
            offset = target - n
            if n in hash and hash[n] != index:
                return [hash[n]+1, index+1]
                break
            hash[offset] = nums.index(n)

        raise Exception()
