from typing import List
import math


class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        def runSim(n: int) -> int:  # return hours it took
            hours = 0
            for p in piles:
                hours += math.ceil(p/n)
            return hours
        maxRate = max(piles)
        res = maxRate
        low, high = 1, maxRate
        while low <= high:
            mid = (low + high) // 2
            midHours = runSim(mid)
            if midHours <= h:
                high = mid - 1
                res = mid
            else:
                low = mid + 1
        return res
