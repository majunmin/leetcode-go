package leetcode_0518

// https://leetcode.cn/problems/coin-change-ii/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	// dp[i] 表示  金钱i 可以有多少种硬币组合数
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
