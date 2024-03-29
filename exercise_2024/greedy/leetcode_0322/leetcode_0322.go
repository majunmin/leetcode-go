package leetcode_0322

// https://leetcode.cn/problems/coin-change/?envType=study-plan-v2&envId=top-100-liked
func coinChange(coins []int, amount int) int {
	cnts := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		cnts[i] = amount + 1
	}
	// init state
	for i := 1; i < amount+1; i++ {
		for _, coin := range coins {
			if i-coin > 0 {
				cnts[i] = min(cnts[i], cnts[i-coins[i]]+1)
			}
		}
	}
	return cnts[amount]
}
