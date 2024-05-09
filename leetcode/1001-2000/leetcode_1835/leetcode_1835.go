package leetcode_1835

// https://leetcode.cn/problems/find-xor-sum-of-all-pairs-bitwise-and/
func getXORSum(arr1 []int, arr2 []int) int {
	// 公式推导： (a&b) ^(a&c) = a&(b^c)
	var num1, num2 int
	for _, num := range arr1 {
		num1 ^= num
	}
	for _, num := range arr2 {
		num2 ^= num
	}
	return num1 & num2
}
