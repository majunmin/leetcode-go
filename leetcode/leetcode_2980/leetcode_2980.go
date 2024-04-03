package leetcode_2980

// https://leetcode.cn/problems/check-if-bitwise-or-has-trailing-zeros/
func hasTrailingZeros(nums []int) bool {
	var cnt int
	for _, num := range nums {
		if num&1 == 0 {
			cnt++
		}
	}
	return cnt >= 2
}
