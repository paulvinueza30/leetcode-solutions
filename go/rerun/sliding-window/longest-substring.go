package slidingwindow

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	res := 0
	set := make(map[byte]bool)
	start, curr := 0, 1
	set[s[start]] = true
	for curr < len(s) {
		if _, ok := set[s[curr]]; ok {
			res = max(res, curr-start)
			for set[s[curr]] {
				delete(set, s[start])
				start++
			}
		}
		set[s[curr]] = true
		curr++
	}

	return max(res, curr-start)
}
