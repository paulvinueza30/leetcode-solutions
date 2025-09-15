package main

import "math"

// greedy fast
func jump(nums []int) int {
	steps := 0
	l, r := 0 , 0
	for r < len(nums) - 1 {
		farthest := 0
		for i := l ; i < (r + 1) ; i++{
			farthest = max(farthest, i + nums[i])
		}
		l = r + 1
		r = farthest
		steps += 1
	}
	return steps
}

// dp not fast

func jumpDP(nums []int) int {
	dp := make([]int, len(nums))
	for i := range len(dp){
		dp[i] = math.MaxInt32
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= len(nums) - 1{
			return 0
		}
		if dp[i] != math.MaxInt32{
			return dp[i]
		}
		for c := 1; c <= nums[i] ; c++ {
			dp[c] = min(dp[c], 1 + dfs(i + c))
		}
		return dp[i]
	}
	return dfs(0)
}