package main

func recursiveClimbStairs(n int) int {
	if n == 0 {
		return 1
	}
	return (n - 1) + (n - 2)
}
func runSimulation(m []int, n int, curr int) int {
	if curr > n {
		return 0
	}
	if curr == n {
		return 1
	}
	if m[curr] != 0 {
		return m[curr]
	}
	m[curr] = runSimulation(m, n, curr+1) + runSimulation(m, n, curr+2)
	return m[curr]
}
func climbStairs(n int) int {
	memo := make([]int, n+1)
	return runSimulation(memo, n, 0)
}
