package main

import "sort"

type Group struct{
	currSize int;
	lastEle int;
}
func isNStraightHand(hand []int, gSize int) bool {
	groups := []Group{}

	sort.Slice(hand, func(i, j int) bool {
		return hand[i] < hand[j]
	})
	for _, c := range hand{
		found := false
		for i := range groups{
			if groups[i].currSize == gSize{
				continue
			}
			if groups[i].lastEle + 1 == c{
				groups[i].lastEle = c
				groups[i].currSize += 1
				found = true
				break
			}
		}
		if !found{
			groups = append(groups, Group{1, c})
		}
	}
	
	for _ , group := range groups{
		if group.currSize != gSize{
			return false 
		 }
		 }
	return true
}