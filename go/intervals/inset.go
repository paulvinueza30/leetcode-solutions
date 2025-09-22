package main

func insert(intervals [][]int, newinterval []int) [][]int {
	var res [][]int

	for i, interval := range intervals{
		if newinterval[1] < interval[0]{
			res = append(res, newinterval)
			return append(res, intervals[i:]...)
		} else if newinterval[0] > interval[1]{
			res = append(res, interval)
		} else {
			newinterval = []int{min(newinterval[0], interval[0]), max(newinterval[1], interval[1])}
		}
	}
	res = append(res, newinterval)
	
	return res
}
