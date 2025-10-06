package main

func findTargetSumWays(nums []int, target int) int {
	dp := make(map[int]int)
	dp[0] = 1

	for _, num := range nums{
		nextDP := make(map[int]int)
		for total , count := range dp {
			nextDP[total + num] += count
			nextDP[total - num] += count
		}
		dp = nextDP
	}
	return dp[target]
}
