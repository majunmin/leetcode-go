package lcr_073

// https://leetcode.cn/problems/nZZqjQ/
func minEatingSpeed(piles []int, h int) int {
	var (
		speed int
	)
	for _, pile := range piles {
		speed = max(speed, pile)
	}
	//二分法
	l, r := -1, speed+1
	for l+1 < r {
		mid := l + (r-l)/2
		if getTime(piles, mid) > h {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

func getTime(piles []int, speed int) int {
	var result int
	for _, pile := range piles {
		result += pile / speed
		if pile%speed != 0 {
			result++
		}
	}
	return result
}
