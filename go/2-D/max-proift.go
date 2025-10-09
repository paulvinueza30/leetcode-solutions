package main

func maxProfit(prices []int) int {
	var dfs func(i int , holding bool , lastBought int) int 
	dfs = func(i int, holding bool, lastBought int) int {
		if i == len(prices) {
			return 0
		}
		skip := dfs(i + 1, holding, lastBought)
		if holding {
			sell := prices[i] - lastBought + dfs(i + 1, false, 0)
			return max(sell , skip)
		} else {
			buy := dfs(i + 1, true, prices[i])
			return max(buy, skip)
		}
	}
	
	return dfs(0 , true, 0)

}

func main(){
	println(maxProfit([]int{1,3,4,0,4}))
}