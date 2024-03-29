package leetcode_0518

// https://leetcode.cn/problems/coin-change-ii/
func change(amount int, coins []int) int {
	length := len(coins)
	if amount == 0 {
		return 1
	}
	if length == 0 {
		return 0
	}
	// 状态定义
	// dp[i][j] : coins[i] 组成金额 j 的最小方案数
	//
	dp := make([][]int, length)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}
	// init state
	// 初始化第一行
	dp[0][0] = 1
	for i := 1; i*coins[0] <= amount; i++ {
		dp[0][i*coins[0]] = 1
	}
	// 动态规划
	for i := 1; i < length; i++ {
		for j := 0; j <= amount; j++ {
			// dp[i][j] = dp[i-1][j] if k == 0
			for k := 0; k*coins[i] < j; k++ {
				dp[i][j] += dp[i-1][j-k*coins[i]]
			}
		}
	}
	return dp[length-1][amount]
}

// 利用 01 背包的 优化算法
func packSolution(amount int, coins []int) int {
	dp := make([]int, amount+1)
	//init state
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
