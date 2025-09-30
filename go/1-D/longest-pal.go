package main

func longestPalindrome(s string) string {
	var res string

	var palCheck func(l, r int)
	palCheck = func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r]{
			if len(s[l:r+1]) > len(res){
				res = s[l:r+1]
			}
			l--
			r++
		}
	}
	for i := range s{
		palCheck(i , i)
		palCheck(i , i + 1)
	}
	return res
}

