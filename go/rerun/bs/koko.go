package bs

func minEatingSpeed(piles []int, h int) int {
	helper := func(k int) int {
		currHour := 0
		for _, p := range piles {
			currHour += (k + p - 1) / k
		}
		return currHour
	}

	low, high := 1, 0
	for _, p := range piles {
		if p > high {
			high = p
		}
	}
	res := high
	for low <= high {
		mid := (low + high) / 2
		if helper(mid) > h {
			low = mid + 1 // eating to slow
		} else {
			res = mid
			high = mid - 1 // good keep trying lower k
		}
	}

	return res
}
