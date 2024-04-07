package leetcode_2871

// https://leetcode.cn/problems/split-array-into-maximum-number-of-subarrays/
func maxSubarrays(nums []int) int {
	// param check
	if len(nums) < 2 {
		return len(nums)
	}
	var (
		result int
		a      = -1 // 全1
	)
	// 1. num 越与越小
	// 2. 子数组之和 尽可能小, 子数组越少越小.(除非各个子数组的分数都是 0) // 子数组的 AND，不会低于整个 nums 数组的 AND。
	for _, num := range nums {
		a &= num
		if a == 0 {
			result += 1
			a = -1
		}
	}
	// 如果 result == 0 , 说明 所有数相与 > 0
	return result
}
