package leetcode_1829

// https://leetcode.cn/problems/maximum-xor-for-each-query/
func getMaximumXor(nums []int, maximumBit int) []int {
	// param check
	if len(nums) == 0 || maximumBit == 0 {
		return nil
	}
	var (
		mask   = 1<<maximumBit - 1
		result = make([]int, len(nums))
	)
	preXor := make([]int, len(nums)+1)
	for i, num := range nums {
		preXor[i+1] = preXor[i] ^ num
		//result[len(nums)-i-1] = maxVal - preXor[i+1]
		result[len(nums)-i-1] = mask ^ preXor[i+1]
	}
	return result
}
