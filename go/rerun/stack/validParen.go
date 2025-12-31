package stack

func isValid(s string) bool {
	MATCHES := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	leftBracketStack := make([]rune, 0)

	for _, c := range s {
		// left bracket add to stack
		if _, ok := MATCHES[c]; ok {
			leftBracketStack = append(leftBracketStack, c)
			continue
		} else if len(leftBracketStack) == 0 {
			return false
		}
		// right bracket
		top := leftBracketStack[len(leftBracketStack)-1]
		topMatching := MATCHES[top]
		if topMatching != c {
			return false
		}
		leftBracketStack = leftBracketStack[:len(leftBracketStack)-1]
	}

	if len(leftBracketStack) != 0 {
		return false
	}
	return true
}
