package leetcode_2683

// https://leetcode.cn/problems/neighboring-bitwise-xor/
func doesValidArrayExist(derived []int) bool {
	var result int
	for _, num := range derived {
		result ^= num
	}
	return result == 0
}
