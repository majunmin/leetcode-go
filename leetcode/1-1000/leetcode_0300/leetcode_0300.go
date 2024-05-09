package leetcode_0300

// https://leetcode.cn/problems/longest-increasing-subsequence/?envType=study-plan-v2&envId=top-100-liked
func lengthOfLIS(nums []int) int {
	// param check
	if len(nums) < 2 {
		return 1
	}
	// 状态定义
	// dp[i] : nums[:i+1] 的最长子序列
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	// 递推公式
	// dp[i] = max(dp[i], nums[])
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	var result int
	for i := 0; i < len(dp); i++ {
		result = max(result, dp[i])
	}
	return result
}
