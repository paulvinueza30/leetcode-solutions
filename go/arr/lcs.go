package main

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}
	maxSeq := 0
	for _, num := range nums {
		if _, ok := set[num-1]; !ok {
			currSeq := 0
			for set[num] {
				currSeq++
				num++
			}
			maxSeq = max(maxSeq, currSeq)
		}
	}
	return maxSeq
}
