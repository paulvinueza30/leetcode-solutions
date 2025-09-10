package main

func reverse(x int) int {
	MAX := (1 << 31) - 1
	MIN := -(1 << 31)
	res := 0
	for x != 0 {
		rightMost := x % 10
		x = int(x / 10)	
		if res > int(MAX / 10) || res < int(MIN / 10){
            return 0
        }
		res = (res * 10) + rightMost
	}
	return res
}