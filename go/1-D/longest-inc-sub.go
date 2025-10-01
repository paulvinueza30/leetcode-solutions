package main

// wont work bc we just take the next greasest number
// func lengthOfLIS(nums []int) int {
// 	var res int
// 	seq := make(map[int][]int)
// 	for i , num := range nums{
// 		found := false
// 		for _, arr := range seq{
// 			if arr[len(arr) - 1] < num{
// 				arr = append(arr, num)
// 				found = true
// 				res = max(res , len(arr))
// 			}
// 		}
// 		if !found {
// 			seq[i] = []int{num}
// 		}
// 	}
// 	return res
// }
//

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = 1
	}
	res := 1

	for i := len(nums) - 1 ; i >= 0 ; i--{
		for j := i + 1 ; j < len(nums) ; j++{
			if nums[i] < nums[j]{
				dp[i] = max(dp[i], 1 + dp[j])
				res = max(res, dp[i])
			}
		}
	}
	return res
}