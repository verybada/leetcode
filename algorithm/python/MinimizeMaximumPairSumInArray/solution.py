class Solution:
    def minPairSum(self, nums: list[int]) -> int:
        nums.sort()

        max_value = 0
        nums_len = len(nums)
        for i in range(nums_len // 2):
            max_value = max(nums[i] + nums[nums_len - 1 - i], max_value)
        return max_value
