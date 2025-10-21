from typing import List


class Solution:
    def longestConsecutive(self, nums: List[int]):
        res = 0
        numSet = set(nums)
        for num in numSet:
            if num - 1 not in numSet:
                streak = 1
                while num + 1 in numSet:
                    num += 1
                    streak += 1
                res = max(res, streak)
        return res
