package backtrack

import "unicode"

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	numToLetters := map[rune]string{
		'2': "ABC",
		'3': "DEF",
		'4': "GHI",
		'5': "JKL",
		'6': "MNO",
		'7': "PRQS",
		'8': "TUV",
		'9': "WXYZ",
	}
	res := make([]string, 0)
	var backtrack func(digitsIdx int, curr string)
	backtrack = func(digitsIdx int, curr string) {
		if digitsIdx == len(digits) {
			res = append(res, curr)
			return
		}
		num := rune(digits[digitsIdx])
		letters := numToLetters[num]
		for _, c := range letters {
			undercase := unicode.ToLower(c)
			backtrack(digitsIdx+1, curr+string(undercase))
		}
	}
	backtrack(0, "")
	return res
}
