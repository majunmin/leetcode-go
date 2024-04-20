package leetcode_0377

// https://leetcode.cn/problems/combination-sum-iv/
func combinationSum4(nums []int, target int) int {
	var (
		dp = make([]int, target+1)
	)
	// init state
	dp[0] = 0
	for _, num := range nums {
		if num <= target {
			dp[num] = 1
		}
	}
	for i := 1; i < len(dp); i++ {
		for _, num := range nums {
			if i-num > 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[len(dp)-1]
}
