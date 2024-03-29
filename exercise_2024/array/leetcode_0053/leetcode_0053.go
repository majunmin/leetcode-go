package leetcode_0053

// https://leetcode.cn/problems/maximum-subarray/?envType=study-plan-v2&envId=top-100-liked
func maxSubArray(nums []int) int {
	// param check
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums)+1)
	//
	for i := 1; i <= len(nums); i++ {
		dp[i] = max(0, dp[i-1]+nums[i-1], nums[i-1])
	}
	return dp[len(dp)-1]
}
