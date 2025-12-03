package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	suf := make([]int, len(nums))
	pre := make([]int, len(nums))

	pre[0] = 1
	suf[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}

	for i := range nums {
		res[i] = pre[i] * suf[i]
	}
	return res
}
