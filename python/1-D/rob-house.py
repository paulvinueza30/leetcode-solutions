from typing import List


class Solution:
    def rob(self, nums: List[int]) -> int:
        dp = [-1] * len(nums)

        def runSim(i: int) -> int:
            if i >= len(nums):
                return 0
            if dp[i] != -1:
                return dp[i]
            robThisHouse = nums[i] + runSim(i + 2)
            skipHouse = runSim(i + 1)
            dp[i] = max(robThisHouse, skipHouse)
            return dp[i]
        return runSim(0, 0)
