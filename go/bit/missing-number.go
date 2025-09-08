package main

func missingNumber(nums []int) int {
	res := len(nums) * ((len(nums) + 1) / 2)
	for _, n := range nums{
		res -= n
	}
	return res
}
