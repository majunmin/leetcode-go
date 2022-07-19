package leetcode_0231

//https://leetcode.cn/problems/power-of-two/
func isPowerOfTwo(n int) bool {
	// 1. if hammingWeight == 1
	return n > 0 && n&(n-1) == 0
}
