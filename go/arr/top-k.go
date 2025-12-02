package main

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	var res []int

	for _, num := range nums {
		counts[num] += 1
	}
	for range k {
		var max int
		var maxNum int
		for k, v := range counts {
			if v >= max {
				max = v
				maxNum = k
			}
		}
		res = append(res, maxNum)
		delete(counts, maxNum)
	}
	return res
}
