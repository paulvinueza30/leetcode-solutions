package main


func reverseBits(n int) int {
	res := 0
	for range 32 {
		res <<= 1
		lastBit := n & 1
		n >>= 1
		res |= lastBit
	}
	return res
}
