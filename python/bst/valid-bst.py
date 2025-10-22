from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        if root is None:
            return True

        def dfs(left: float, right: float, curr: Optional[TreeNode]) -> bool:
            if curr is None:
                return True
            if not left < curr.val < right:
                return False
            return dfs(curr.val, right, curr.right) and dfs(left, curr.val, curr.left)
        return dfs(float("-inf"), float("inf"), root)
