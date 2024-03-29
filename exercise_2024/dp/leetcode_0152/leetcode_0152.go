package leetcode_0152

// https://leetcode.cn/problems/maximum-product-subarray/?envType=study-plan-v2&envId=top-100-liked
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		panic("invalid param")
	}
	// 状态定义
	//dp[i][0]:以 i 为结尾的 乘积最小子数组
	//dp[i][1]:以 i 为结尾的 乘积最大子数组
	// init state
	dp := make([][2]int, len(nums))
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]
	for i := 1; i < len(dp); i++ {
		if nums[i] >= 0 {
			dp[i][0] = min(nums[i], dp[i-1][0]*nums[i])
			dp[i][1] = max(nums[i], dp[i-1][1]*nums[i])
		} else {
			dp[i][0] = min(nums[i], dp[i-1][1]*nums[i])
			dp[i][1] = max(nums[i], dp[i-1][0]*nums[i])
		}
	}

	result := dp[0][1]
	for i := 1; i < len(dp); i++ {
		result = max(result, dp[i][1])
	}
	return result
}
