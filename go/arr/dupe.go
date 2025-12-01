package main

func hasDuplicate(nums []int) bool {
	numSet := make(map[int]bool, 0)
	for _, num := range nums {
		if numSet[num] {
			return true
		}
		numSet[num] = true
	}
	return false
}
