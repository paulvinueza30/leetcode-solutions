package main

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum % 2 != 0 {
		return false
	}
	target := int(sum / 2)
	dp := make(map[int]bool)
	dp[0] = true
	for i := len(nums) - 1 ; i >= 0 ; i-- {
		nextDP := make(map[int]bool)
		for k, _ := range dp {
			nextDP[k + nums[i]] = true
			nextDP[k] = true
		}
		dp = nextDP
	}
	_, ok := dp[target]
	return ok
}