class Solution:
    def smallestEqual(self, nums: List[int]) -> int:
        smaller_index = -1
        for index in range(len(nums)-1, -1, -1):
            num = nums[index]
            if (index % 10) != num:
                continue
                
            smaller_index = index
        return smaller_index  