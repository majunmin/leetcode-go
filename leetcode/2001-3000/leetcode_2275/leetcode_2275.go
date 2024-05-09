package leetcode_2275

// https://leetcode.cn/problems/largest-combination-with-bitwise-and-greater-than-zero/description/
func largestCombination(candidates []int) int {
	var result int
	// int 32 bit,如果组合中所有数字的在同一个bit位都是1, 那么 这组数的 & > 0,
	// 可以按bit 位统计.
	for i := 0; i < 32; i++ {
		var cnt int
		for _, num := range candidates {
			if num>>i&1 == 1 {
				cnt++
			}
		}
		result = max(result, cnt)
	}
	return result
}
