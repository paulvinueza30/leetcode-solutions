package main

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
		n *= -1
	}

}

func powHelper(x float64, n int) float64{
	
}