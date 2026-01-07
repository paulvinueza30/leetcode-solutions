package twopointers

func reverseVowels(s string) string {
	l, r := 0, len(s)-1

	vowelMap := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	isVowel := func(c rune) bool {
		if _, ok := vowelMap[c]; ok {
			return true
		}
		return false
	}
	rev := []rune(s)
	for l < r {
		for l < r && !isVowel(rune(s[l])) {
			l++
		}
		for l < r && !isVowel(rune(s[r])) {
			r--
		}
		tmp := rev[l]
		rev[l] = rev[r]
		rev[r] = tmp
		l++
		r--
	}

	return string(rev)
}
