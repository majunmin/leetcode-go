package exercise_2024

func maxProduct(nums []int) int {
	// param check
	dp := make([][2]int, len(nums)+1)
	// define
	//dp[i][0] 以第 i-1 位结尾的 最大乘积 (>0)
	//dp[i][1] 以第 i-1 位结尾的 的最小乘积(< 0)
	// init state

	var result int
	// 状态转移

	for i := 1; i < len(dp); i++ {
		if nums[i-1] >= 0 {
			dp[i][0] = max(nums[i-1], dp[i-1][0]*nums[i])
			dp[i][1] = min(0, dp[i-1][1]*nums[i])
		} else {
			dp[i][0] = max(0, dp[i-1][1]*nums[i])
			dp[i][1] = min(0, dp[i-1][0]*nums[i])
		}
		result = max(result, dp[i][0])
	}
	return result
}
