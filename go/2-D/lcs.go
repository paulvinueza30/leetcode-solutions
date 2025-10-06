package main

func longestCommonSubsequence(s1 , s2 string) int {
	dp := make([][]int, len(s1))
	for i := range s1 {
		dp[i] = make([]int, len(s2))
		for j := range s2 {
			dp[i][j] = -1
		}
	}
	var dfs func(i, j int) int 
		dfs = func(i, j int) int {
		if i == len(s1) || j == len(s2) {
			return 0
		}
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		if s1[i] == s2[j] {
			dp[i][j] = 1 + dfs(i + 1, j + 1)
		} else {
			dp[i][j] = max(dfs(i + 1, j) , dfs(i , j + 1))
		}
		return dp[i][j]
	}
	return dfs(0, 0)
}