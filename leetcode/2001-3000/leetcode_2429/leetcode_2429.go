package leetcode_2429

import "math/bits"

// https://leetcode.cn/problems/minimize-xor/
func minimizeXor(num1 int, num2 int) int {
	//题解: https://leetcode.cn/problems/minimize-xor/description/
	// 贪心的思想
	c1, c2 := bits.OnesCount(uint(num1)), bits.OnesCount(uint(num2))
	for ; c1 < c2; c2-- {
		num1 ^= num1 - 1 // 将最低位的1置为 0
	}
	for ; c1 > c2; c2-- {
		num1 |= num1 + 1 // 将最低位的 0 置为1
	}
	// 当 c1 == c2 时, target == num1 结果最小
	return num1
}
