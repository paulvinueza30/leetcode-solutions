package main

func singleNumber(nums []int) int {
	res := nums[0]
	for _ , n := range nums[1:]{
		res ^= n
	}
	return res
}