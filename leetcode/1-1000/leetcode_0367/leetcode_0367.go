package leetcode_0367

// https://leetcode.cn/problems/valid-perfect-square/
func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left < right {
		mid := left + (right-left+1)/2
		if mid*mid > num {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left*left == num
}
