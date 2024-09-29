package leetcode_2740

import "sort"

// https://leetcode.cn/problems/find-the-value-of-the-partition/?envType=daily-question&envId=2024-07-26
func findValueOfPartition(nums []int) int {
	if len(nums) <= 1 {
		// 参数异常
		panic("invalid param, nums.len should greater than 1.")
	}
	sort.Ints(nums)
	res := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		res = min(res, nums[i]-nums[i-1])
	}
	return res
}
