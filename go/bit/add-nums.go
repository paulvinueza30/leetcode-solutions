package main

func getSum(a int, b int) int {
	res := []int{}
	carryOver := 0
	for i := 0 ; i < 32 ; i++ {
		aBit , bBit := a & 1 , b & 1
		if (aBit & bBit & carryOver) == 1{
			res = append(res, 1)
			carryOver = 1
		} else if (aBit & bBit) == 1 || (aBit & carryOver) == 1 || (bBit & carryOver) == 1 {
			res = append(res, 0)
			carryOver = 1
		} else {
			res = append(res, aBit | bBit | carryOver) 
			carryOver = 0
		}
		a >>= 1
		b >>= 1
	}
	final := 0
	for i := len(res) - 1; i >= 0 ; i--{
		final <<= 1
		final |= res[i]
	}
	return int(int32(final))
	}