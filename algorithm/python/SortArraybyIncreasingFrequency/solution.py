from collections import Counter


class Solution:
    def frequencySort(self, nums: list[int]) -> list[int]:
        counter = Counter(nums)
        nums.reverse()
        return sorted(nums, key=lambda x: (counter[x], -x))
