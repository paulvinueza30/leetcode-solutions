package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int][]int)
	for i, n := range nums {
		numMap[n] = append(numMap[n], i)
	}
	for i, n := range nums {
		v := target - n
		vSlice := numMap[v]
		if n == v && len(vSlice) > 1 {
			return []int{i, vSlice[len(vSlice)-1]}
		}
		if vSlice != nil {
			return []int{i, vSlice[len(vSlice)-1]}
		}
	}
	return []int{0, 0}
}
