package main

func sim(cost []int, memo []int, idx int) int {
	if idx >= len(cost) {
		return 0
	}
	if memo[idx] != -1 {
		return memo[idx]
	}

	memo[idx] = cost[idx] + min(sim(cost, memo, idx+1), sim(cost, memo, idx+2))
	return memo[idx]
}

func minCostClimbingStairs(cost []int) int {
	m := make([]int, len(cost))
	for i := range len(cost) {
		m[i] = -1
	}
	return min(sim(cost, m, 0), sim(cost, m, 1))

}
