package lintcode_0563

// https://www.lintcode.com/problem/563/?showListFe=true&page=1&submissionStatus=ACCEPTED&pageSize=50
func BackPackV(nums []int, target int) int {
	// write your code here
	// param check
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, target+1)
	// init state
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := nums[i]; j <= target; j++ {
			if dp[j-nums[i]] > 0 {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
