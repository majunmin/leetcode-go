package leetcode_2595

import "math/bits"

// https://leetcode.cn/problems/number-of-even-and-odd-bits/
func evenOddBit(n int) []int {
	result := make([]int, 2)
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			result[i&1]++
		}
		n >>= 1
	}
	return result
}

func evenOddBit2(n int) []int {
	// 构造掩码
	mask := 0x5555
	return []int{bits.OnesCount(uint(n & (mask >> 1))), bits.OnesCount(uint(n & mask))}
}
