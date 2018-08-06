class Solution:
    def majorityElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        middle = len(nums) // 2
        nums.sort()
        return nums[middle]
    
        # hard work solution
        maps = dict()
        repeat_times = 0
        num = None
        for n in nums:
            if not n in maps:
                maps[n] = 0
            maps[n] += 1
            
            if num == n:
                repeat_times += 1
                continue
                
            n_repeated = maps[n]
            if n_repeated > repeat_times:
                repeat_times = n_repeated
                num = n
        return num
        
            
