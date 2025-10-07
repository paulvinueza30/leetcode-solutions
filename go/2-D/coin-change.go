package main

import "fmt"

// exponential solution
// func change(amount int, coins []int) int {
// 	var dfs func(i int, curr int) int
// 	dfs = func(i int, curr int) int{
// 		if curr == amount{
// 			return 1
// 		}
// 		if i == len(coins) || curr > amount{
// 			return 0
// 		}
// 		return dfs(i , curr + coins[i]) + dfs(i + 1, curr)
// 	}

// 	return dfs(0, 0)
// }
func change(amount int , coins []int) int{
	dp := make(map[[2]int]int)
	var dfs func(i , curr int) int 
	dfs = func(i , curr int) int{
		if curr == amount {
			return 1
		}
		if i == len(coins) || curr > amount{
			return 0
		}
		cache := [2]int{i, curr}
		if v , ok := dp[cache]; ok{
			return v
		}
		dp[cache] = dfs(i , curr + coins[i]) + dfs(i + 1, curr)
		return dp[cache]
	}
	return dfs(0 , 0)
	}

func change2(amount int, coins []int) int {
	dp := make([]int, amount + 1)
	dp[0] = 1
	
	for i := len(coins) - 1 ; i >= 0 ; i-- {
		nextDP := make([]int , amount + 1)
		nextDP[0] = 1
		for a := 1 ; a <= amount ; a++{
			nextDP[a] = dp[a]
			if a - coins[i] >= 0{
				nextDP[a] += nextDP[a - coins[i]]
			}
		}
		dp = nextDP
	}
	return dp[amount]
}

func main(){
	fmt.Println(change2(4, []int{1,2,3}))
}