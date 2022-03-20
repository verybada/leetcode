from typing import List
import heapq


class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        if k == len(points):
            return points

        results = list()
        for x, y in points:
            distance = (x*x) + (y*y)
            results.append((distance, [x, y]))

        heapq.heapify(results)
        return [point for _, point in heapq.nsmallest(k, results)]
