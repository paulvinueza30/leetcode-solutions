package twopointers

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l <= r {
		if !isAlpha(s[l]) {
			l++
			continue
		} else if !isAlpha(s[r]) {
			r--
			continue
		}
		if !equalIgnoreCase(s[l], s[r]) {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlpha(c byte) bool {
	if '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' {
		return true
	}
	return false
}

func equalIgnoreCase(a, b byte) bool {
	if a == b {
		return true
	}
	if a >= 'A' && a <= 'Z' {
		a += 32
	}
	if b >= 'A' && b <= 'Z' {
		b += 32
	}
	return a == b
}
