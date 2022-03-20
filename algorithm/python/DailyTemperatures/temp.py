from typing import List


class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        results = [0] * len(temperatures)

        past_daily_temps = list()
        for day, temp in enumerate(temperatures):
            while (past_daily_temps and temp > past_daily_temps[-1][1]):
                past_day, past_temp = past_daily_temps.pop()
                results[past_day] = day - past_day
            past_daily_temps.append([day, temp])
        return results
