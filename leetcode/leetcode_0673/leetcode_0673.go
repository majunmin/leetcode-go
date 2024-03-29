package leetcode_0673

// https://leetcode.cn/problems/number-of-longest-increasing-subsequence/
func findNumberOfLIS(nums []int) int {
	// param check
	if len(nums) < 2 {
		return len(nums)
	}
	// process
	var (
		dp        = make([]int, len(nums))
		cnts      = make([]int, len(nums))
		maxLength int
		result    int
	)
	// 状态定义
	// cnts[i] 以 i 为结尾的 最长子数组的个数
	dp[0] = 1
	cnts[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// 第一次找到
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnts[i] = cnts[j]
				} else if dp[j]+1 == dp[i] {
					// 再次找到
					cnts[i] += cnts[j]
				}
			}
		}
		maxLength = max(maxLength, dp[i])
	}
	for i, length := range dp {
		if length == maxLength {
			result += cnts[i]
		}
	}
	return result
}
