package lcr_104

// 1 2 3
// 0 1
// 1 1
// 2 11 2
// 3 111 21 12 3
// 4 1111 112 121 13 211 31 6

// https://leetcode.cn/problems/D0F0SV/
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	// init state
	dp[0] = 1
	// 状态转移
	//dp[i] += dp[i - (nums...)]
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func backtrace(nums []int, target int) int {
	// terminate
	if target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}
	// for choice in choiceList
	var cnt int
	for i := 0; i < len(nums); i++ {
		cnt += backtrace(nums, target-nums[i])
	}
	return cnt
}
