package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	for i := range dp{
		dp[i] = amount + 1
	}
	dp[0] = 0
	for a := range (amount + 1){
		for _, c := range coins{
			if a - c >= 0{
				dp[a] = min(dp[a] , 1 + dp[a-c]) 
			}
		}
	}
	if dp[amount] == amount + 1{
		return -1
	}
	return dp[amount]
}