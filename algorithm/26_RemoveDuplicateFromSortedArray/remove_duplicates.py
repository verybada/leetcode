class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if not nums:
            return 0

        tail = 0
        for i in range(1, len(nums)):
            if nums[i] != nums[tail]:
                tail += 1
                nums[tail] = nums[i]

        return tail + 1


if __name__ == '__main__':
    nums = []
    assert Solution().removeDuplicates(nums) == 0 and nums == []
    nums = [1]
    assert Solution().removeDuplicates(nums) == 1 and nums == [1]
    nums = [1, 2]
    assert Solution().removeDuplicates(nums) == 2 and nums == [1, 2]
    nums = [1, 1, 2]
    assert Solution().removeDuplicates(nums) == 2 and nums == [1, 2, 2]
    nums = [1, 1, 2, 2, 3, 3]
    assert Solution().removeDuplicates(nums) == 3 and nums == [1, 2, 3, 2, 3, 3]
