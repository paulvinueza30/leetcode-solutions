package bs

func search(nums []int, target int) int {
	l, h := 0, len(nums)-1

	for l <= h {
		mid := (l + h) / 2
		v := nums[mid]
		switch {
		case v == target:
			return mid
		case v < target:
			l = mid + 1
		case v > target:
			h = mid - 1
		}
	}
	return -1
}
