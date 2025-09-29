package main

import "sort"

func minMeetingRooms(intervals []interval) int {
	start , end := []int{}, []int{}
	res , count := 0 , 0
	for _ , interval := range intervals{
		start = append(start, interval.start)
		end = append(end, interval.end)
	}
	sort.Slice(start , func(i, j int) bool {
		return start[i] < start[j]
	})
	sort.Slice(end, func(i, j int) bool {
		return end[i] < end[j]
	})
	sp , ep := 0 , 0

	for sp < len(start){
		if start[sp] < end[ep] {
			count++
			sp++
			res = max(res , count)
		} else{
			count--
			ep++
		}
	}
	return res
}
