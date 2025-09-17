package main

// "xyxxyzbzbbisl"

// slower but works
// func partitionLabels(s string) []int {
// 	res := []int{}
// 	findLastIdx := func(c byte) int {
// 		for i := len(s) - 1; i >= 0 ; i--{
// 			if s[i] == byte(c){
// 				return i
// 			}
// 		}
// 		return -1
// 	}
// 	l := 0
// 	seen := make(map[byte]bool)
// 	for l < len(s) {
// 		seen[s[l]] = true
// 		lastIdx := findLastIdx(s[l])
// 		curr := l
// 		for curr < lastIdx{
// 			_, v := seen[s[curr]]
// 			if !v{
// 				newIdx := findLastIdx(s[curr])
// 				if newIdx > lastIdx{
// 					lastIdx = newIdx
// 				}
// 			}
// 			curr++
// 		}
// 		res = append(res, (lastIdx-l+1))
// 		l = lastIdx + 1
// 	}

// 	return res
// }

func partitionLabels(s string) []int{
	res := []int{}

	lastOcc := make(map[rune]int)
	for i , c := range s{
		lastOcc[c] = i
	}
	
	start, lastIdx := 0, 0
	for i , c := range s {
		lastIdx = max(lastIdx, lastOcc[c])
		if i == lastIdx{
			res = append(res, (lastIdx - start))
			start = lastIdx + 1
		}
	}
	return res
}