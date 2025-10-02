package main

// Time : O(n^2) Space : O(n)
// func isPal(s string) bool {
// 	l , r := 0 , len(s) - 1
// 	for l < r {
// 		if s[l] != s[r]{
// 			return false
// 		}
// 		l++
// 		r--
// 	}
// 	return true
// }
// func countSubstrings(s string) int {
// 	seen := make(map[string]bool)
// 	res := 0
// 	for l := range s {
// 		seen[string(s[l])] = true
// 		for r := l + 1 ; r < len(s) ; r++ {
// 			if seen[s[l:r+1]]{
// 				res++
// 				continue
// 			}
// 			if isPal(s[l : r + 1]){
// 				seen[s[l : r + 1]] = true
// 				res++
// 			}
// 		}
// 	}
// 	return res
// }

func countSubstrings(s string) int {
	countPalis := func(l , r int) int {
		res := 0
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}
		return res
	}
	res := 0
	for i := range s {
		res += countPalis(i , i) + countPalis(i , i + 1)
	}
	return res
}