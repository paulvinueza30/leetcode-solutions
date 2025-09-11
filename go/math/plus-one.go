package main

// only works if nums are 32 bits long
// func plusOne(digits []int) []int {
// 	num := 0
// 	for _, n := range digits{
// 		num *= 10
// 		num += n
// 	}
// 	num++
// 	idx := len(digits) - 1
// 	for num > 0 {
// 		if idx == -1 {
// 			firstElem := []int{num % 10}
// 			digits = append(firstElem, digits...)
// 			break
// 		}
// 		digits[idx] = num % 10
// 		num = int(num/10)
// 		idx--
// 	}
// 	return digits
// }
func plusOne(digits []int) []int {
	carry := 0
	for i := len(digits) - 1 ; i >= 0 ; i-- {
		curr := digits[i] + 1 + carry
		if curr > 9 {
			curr = 0
			carry = 1
		} else {
			return digits
		}
		digits[i] = curr
	}
	if carry == 1{
		digits = append([]int{1}, digits...)
	}
	return digits
}
