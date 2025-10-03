package main

func uniquePaths(m , n int) int{
	dp := make([][]int, m) 
	for i := range m {
		dp[i] = make([]int, n)
	}
	var dfs func(r , c int) int
	dfs = func(r , c int) int{
		if (r >= m || r < 0 || c >= n || c < 0){
			return 0
		}
		if r == (m - 1) && c == (n -1) {
			return 1
		}
		if dp[r][c] != 0 {
			return dp[r][c]
		}
		dp[r][c] = dfs(r + 1, c) + dfs(r , c + 1)
		return dp[r][c]
	}

	return dfs(0, 0)
}