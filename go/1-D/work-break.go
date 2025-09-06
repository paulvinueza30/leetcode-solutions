package main

func wordBreak(s string, wordDict []string) bool {
	set := make(map[string]bool)
	for _, w := range wordDict {
		set[w] = true
	}
	dp := make([]bool, len(s) + 1)
	dp[0] = true
	
	for r := range(len(s) + 1){
		for l := range(r) {
			if dp[l] && set[s[l:r]] {
				dp[r] = true
				break
			}
		}
	}
	return dp[len(s)]
}