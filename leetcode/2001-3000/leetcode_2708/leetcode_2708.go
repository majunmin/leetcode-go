package leetcode_2708

// https://leetcode.cn/problems/maximum-strength-of-a-group/
func maxStrength(nums []int) int64 {
	// 动态规划
	var (
		mx = nums[0]
		mn = nums[0]
	)
	for i := 1; i < len(nums); i++ {
		mx, mn = max(mx, nums[i], mn*nums[i], mx*nums[i]), min(mn, nums[i], mn*nums[i], mx*nums[i])
	}
	return int64(mx)
}
