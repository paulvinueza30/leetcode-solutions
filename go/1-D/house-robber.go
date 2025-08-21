package main

func robSim(n []int, m []int, curr int) int {
	if curr >= len(n) {
		return 0
	}
	if m[curr] != -1 {
		return m[curr]
	}
	m[curr] = max(n[curr]+robSim(n, m, curr+2), robSim(n, m, curr+1))
	return m[curr]
}

func rob(nums []int) int {
	m := make([]int, len(nums))

	for i := range len(nums) {
		m[i] = -1
	}
	return robSim(nums, m, 0)
}
