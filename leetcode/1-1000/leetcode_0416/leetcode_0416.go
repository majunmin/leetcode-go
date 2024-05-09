package leetcode_0416

// https://leetcode.cn/problems/partition-equal-subset-sum
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 { // 奇数 无法拆分
		return false
	}
	target := sum >> 1

	// 定义状态
	//dp[i][j] 对于前i个物品,当前背包容量为 j 时,是否可以将背包装满
	// dp[i][j] = dp[i-1][j - wt[i-1]](装入背包) dp[i-1][j] (不装入背包)
	//
	dp := make([][]bool, 0, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp = append(dp, make([]bool, target+1))
	}

	// init state
	for i := 0; i < target+1; i++ {
		dp[i][0] = true
	}

	// 状态转移
	for i := 1; i < len(nums)+1; i++ {
		for j := 1; j < target+1; j++ {
			if j-nums[i-1] < 0 { // 当前物品装不下时, 无法装下 第i个物品
				dp[i][j] = dp[i-1][j]
			} else { // 可以装下时 可以选择装入或者不装入背包
				dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][target]
}
