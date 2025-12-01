package main

//	func isAnagram(s string, t string) bool {
//		if len(s) != len(t) {
//			return false
//		}
//		sSet := make(map[rune]int)
//		tSet := make(map[rune]int)
//		for i := range s {
//			var ok bool
//			var sSetCnt int
//			var tSetCnt int
//			if sSetCnt, ok = sSet[rune(s[i])]; !ok {
//				sSetCnt = 0
//			}
//			if tSetCnt, ok = tSet[rune(t[i])]; !ok {
//				tSetCnt = 0
//			}
//
//			sSet[rune(s[i])] = sSetCnt + 1
//			tSet[rune(t[i])] = tSetCnt + 1
//		}
//		for k := range sSet {
//			if _, ok := tSet[k]; !ok {
//				return false
//			}
//			if tSet[k] != sSet[k] {
//				return false
//			}
//		}
//		return true
//	}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sSet := make(map[rune]int)
	tSet := make(map[rune]int)

	for _, c := range s {
		sSet[c] += 1
	}
	for _, c := range t {
		tSet[c] += 1
	}

	for k := range sSet {
		if sSet[k] != tSet[k] {
			return false
		}
	}
	return true
}
