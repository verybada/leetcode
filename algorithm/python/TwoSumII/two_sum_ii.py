from typing import List


class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        results = dict()
        for index, number in enumerate(numbers):
            remain = target - number
            previous_index = results.get(remain)
            if previous_index is not None:
                return [previous_index+1, index+1]
            results[number] = index

        assert 0
