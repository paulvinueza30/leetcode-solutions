from typing import List


class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        # numbers = [1,2,3,4] target = 3 output [1,2]
        l, r = 0, len(numbers) - 1
        while l < r:
            sum = numbers[l] + numbers[r]
            if sum == target:
                return [l + 1, r + 1]
            elif sum > target:
                r -= 1
            else:
                l += 1
