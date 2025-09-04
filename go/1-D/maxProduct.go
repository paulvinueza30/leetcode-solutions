package main

func maxProduct(nums []int) int {
	res := nums[0]
	currMax , currMin := 1, 1
	
	for _, num := range nums{
		if num == 0{
			currMax, currMin = 1, 1
		}
		temp := num * currMax
		currMax = max(temp, num * currMin , num)
		currMin = min(temp, num * currMin, num)
		res = max(res, currMax)
	}
	return res
}