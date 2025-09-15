package main


func getSquareSum(n int) int {
	sum := 0
	for n > 0 {
		rightMost := n % 10
		squared := rightMost * rightMost
		sum += squared
		n = int(n/10)
	}
	return sum
}
func isHappy(n int) bool {
	// hash set slow
	// seen := make(map[int]bool)

	// seen[n] = true
	// for {
	// 	squareSum := getSquareSum(n)
	// 	if squareSum == 1 {
	// 		return true
	// 	}
	// 	_ , ok := seen[squareSum]
	// 	if ok {
	// 		return false
	// 	}
	// 	n = squareSum
	// 	seen[squareSum] = true
	// }
	slow , fast := n , getSquareSum(n)
	for slow != fast {
		slow = getSquareSum(slow)
		fast = getSquareSum(fast)
		fast = getSquareSum(fast)
	}
	return fast == 1
}
