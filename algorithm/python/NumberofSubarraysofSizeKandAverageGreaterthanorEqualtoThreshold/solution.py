class Solution:
    def numOfSubarrays(self, arr: list[int], k: int, threshold: int) -> int:
        count = 0
        sum_ = 0
        threshold *= k
        for index, value in enumerate(arr):
            sum_ += value

            if index < k - 1:
                continue

            if sum_ >= threshold:
                count += 1

            sum_ -= arr[index - k + 1]

        return count
