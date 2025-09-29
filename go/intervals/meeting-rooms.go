package main

import "sort"

type interval struct{
	start int;
	end int;
}
// Input: intervals = [(0,30),(5,10),(15,20)]

func canAttendMeetings(intervals []interval) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})
	for i := 0 ; i < len(intervals) - 1 ; i++ {
		if intervals[i].end > intervals[i + 1].start {
			return false
		}
	}
	return true
}