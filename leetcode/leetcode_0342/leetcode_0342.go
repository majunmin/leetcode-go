package leetcode_0342

// https://leetcode.cn/problems/power-of-four/
func isPowerOfFour(n int) bool {
	// 1   0000001
	// 4   0000100
	// 16  0010000
	// 64  1000000

	// mask = (0101010101010101010101010101010101010101010101010101010101010101)2  // 64bit
	return n > 0 && (n&(n-1)) == 0 && n&0x55555555 == n
}
