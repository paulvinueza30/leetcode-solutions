from typing import List


class Solution:
    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:
        combos = []
        curr = []

        def backtrack(currSum: int, start: int):
            if currSum > target:
                return
            if currSum == target:
                combos.append(curr.copy())
                return
            for i in range(start, len(nums)):
                curr.append(nums[i])
                backtrack(currSum + nums[i], i)
                curr.pop()
        backtrack(0, 0)
        return combos
