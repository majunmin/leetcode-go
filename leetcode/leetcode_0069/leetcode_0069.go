package leetcode_0069

// https://leetcode.cn/problems/sqrtx/
func mySqrt(x int) int {
	// x 的 平方根 的取值范围 [0-x]
	// 求 <= 目标值的第一个值
	left, right := 0, x
	for left < right {
		mid := left + (right-left+1)/2
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	// 退出 循环时 有  left == right 成立
	return left
}
