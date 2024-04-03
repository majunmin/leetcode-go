package leetcode_0693

// https://leetcode.cn/problems/binary-number-with-alternating-bits/description/
func hasAlternatingBits(n int) bool {
	y := n ^ (n >> 1)
	// 如何判断 n ^ y 是全1
	//y&(y-1) 将y 最后一位置 0
	return y&(y-1) == 0
}

func solution1(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	pre := n & 1
	n >>= 1
	for i := 0; n > 0 && i < 32; i++ {
		if 1-n&1 != pre {
			return false
		}
		pre = n & 1
		n >>= 1
	}
	return true
}
