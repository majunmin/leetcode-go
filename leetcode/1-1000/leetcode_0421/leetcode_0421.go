package leetcode_0421

// https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/
// 1. 利用了第一题 两数之和的思路， 找到两数异或结果为target 的两数, 用map.
// 2. 试填法, 从高位到地位 尝试将每一位置为1
// 2.1 计算低位时 如何把高位也考虑进来呢?
// 2.2
func findMaximumXOR(nums []int) int {
	var (
		ans  int
		seen map[int]bool
		mask int
	)
	for i := 31; i >= 0; i-- {
		seen = make(map[int]bool)
		newAns := ans | 1<<i // 1. 尝试将第i位置为1
		mask |= 1 << i
		for _, num := range nums {
			num &= mask
			// 2. 存在  num ^ (num ^ newAns) = ans
			if seen[num^newAns] {
				ans = newAns // 3. 该比特位可以是1
				break
			}
			seen[num] = true
		}
	}
	return ans
}
