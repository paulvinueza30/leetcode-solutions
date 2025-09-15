package main

import "math"

// O(n) too slow for larger n
// func myPow(x float32, n int) float32 {
// 	if n == 0{
// 		return float32(1)
// 	}
// 	val := x
// 	negative := n < 0
// 	if negative {
// 		n *= -1
// 	}
// 	for n != 1{
// 		x *= val
// 		n -= 1
// 	}
// 	if negative{
// 		return float32(1/x)
// 	}
// 	return x
// }

func myPow(x float64, n int) float64{
	if n == 0 || x == 1 {
		return float64(1)
	}
	if n < 0 {
		x = float64(1/x)
		if n == math.MinInt32{
			return x * powHelper(x , -(n+1))
		}
	
		n *= -1
	}

	return powHelper(x , n)
}

func powHelper(x float64, n int) float64{
	if n <= 0 {
		return float64(1)
	}
	half := powHelper(x , n / 2)
	if n % 2 == 0{
		return half * half
	}
	return half * half * x
}