package leetcode_0868

import "math"

// https://leetcode.cn/problems/binary-gap/
// 简单模拟题目 🐶
func binaryGap(n int) int {
	var maxGap int
	end := math.MaxInt
	for i := 0; i < 32; i++ {
		if n>>i&1 == 1 {
			maxGap = max(maxGap, i-end)
			end = i
		}
	}
	return maxGap
}
