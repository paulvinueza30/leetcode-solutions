package main

func numDecodings(s string) int {
	dp := make([]int, len(s))
	for i := range len(dp){
		dp[i] = -1
	}
    var helper func(int) int
	helper = func(curr int) int{
		if curr >= len(s){
			return 1
		}
		if s[curr] == '0'{
			return 0
		}
		if dp[curr] != -1{
			return dp[curr]
		}
		dp[curr] = helper(curr + 1)
		if (curr + 1) < len(s){
            digit := int(s[curr]-'0')*10 + int(s[curr+1]-'0')
            if  10 <= digit && digit <= 26{
			    dp[curr] += helper(curr + 2)
            }
		}
		return dp[curr]
	}
	return helper(0)
}