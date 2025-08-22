package main

func robHouse(n []int, m []int, curr int, start int) int {
	if curr >= len(n) || (curr == len(n) - 1  && start == 0){
		return 0
	}
	if m[curr] != -1{
		return m[curr]
	}
	m[curr] = max(n[curr] + robHouse(n, m , curr + 2, start) , robHouse(n, m , curr + 1, start))
	return m[curr]
}
func rob2(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	m := make([][]int, 2)
	for i := range m{
		m[i] = make([]int, len(nums))
	}
	for i := range len(nums){
		m[0][i] = -1
		m[1][i] = -1
	}
	return max(robHouse(nums, m[0], 0 , 0), robHouse(nums, m[1] , 1, 1))
}

