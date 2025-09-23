package main

import "sort"

func merge(intervals [][]int) [][]int{
	sort.Slice(intervals, func(i int , j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	curr := 0
	for curr < len(intervals) - 1{
		if intervals[curr][1] >= intervals[curr + 1][0]{
			intervals[curr + 1] = []int{min(intervals[curr][0] , intervals[curr + 1][0]) , max(intervals[curr][1], intervals[curr + 1][1])}
			intervals[curr] = []int{-1, -1}
		}

		curr++
	}
	var res [][]int
	for _, interval := range intervals{
		if interval[0] != -1{
			res = append(res, interval)
		}
	}
	return res
}