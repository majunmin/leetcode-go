package leetcode_0633

import "math"

func judgeSquareSum(c int) bool {
	// param check
	if c < 0 {
		return false
	}
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}
	return false
}
