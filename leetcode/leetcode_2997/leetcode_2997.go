package leetcode_2997

import "math/bits"

// https://leetcode.cn/problems/minimum-number-of-operations-to-make-array-xor-equal-to-k/
func minOperations(nums []int, k int) int {
	xors := k
	for _, num := range nums {
		xors ^= num
	}
	return bits.OnesCount(uint(xors))
}
