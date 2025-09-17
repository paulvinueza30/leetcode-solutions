package main

// Input: s = "(((*)"

// Input: s = "((**)"

// type stack []rune

// func (s *stack) Push(c rune) {
// 	*s = append(*s, c)
// }
// func (s *stack) Pop() {
// 	*s = (*s)[:len(*s)-1]
// }
// func (s *stack) isEmpty() bool{
// 	return len(*s) == 0
// }
// func checkValidString(s string) bool{
// 	wildCards := []rune{}
// 	var stak stack
// 	for _ , c := range s{
// 		if c == '('{
// 			stak.Push(c)
// 		} else if c == '*'{
// 			wildCards = append(wildCards, c)
// 		}else{
// 			if stak.isEmpty() && len(wildCards) == 0{
// 				return false
// 			}
// 			if !stak.isEmpty(){
// 				stak.Pop()
// 			} else{
// 				wildCards = (wildCards)[:len(wildCards)-1]
// 			}
// 		}
// 	}
// 	// work work because position matters for wildcards
// 	return len(stak) <= len(wildCards)
// }

func checkValidString(s string) bool {
	minL, maxL := 0 , 0
	for _ , c := range s {
		switch c{
		case '(':
			minL++
			maxL++
		case ')':
			minL--
			maxL--
		case '*':
			minL--
			maxL++
	}
	if maxL < 0 {
		return false
	}
	minL = max(minL, 0)
}
	return minL == 0
}