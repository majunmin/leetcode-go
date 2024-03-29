package knapsack_002

var result int

func maxWeight(w, n int, weight []int, value []int) int {
	dfsSolution(w, n, weight, value, 0, 0, 0)

	//

	//return dpSolution(w, n, weight, value)
	return result
}

func dfsSolution(w int, n int, weight []int, value []int, idx, cw, cv int) {
	// terminate
	if cw == w || idx == n {
		result = max(result, cv)
		return
	}
	// 不选择当前物品
	dfsSolution(w, n, weight, value, idx+1, cw, cv)
	// 选择当前物品
	dfsSolution(w, n, weight, value, idx+1, cw+weight[idx], cv+value[idx])
}

func dpSolution(w int, n int, weight []int, value []int) int {

	// state define
	// dp[i][j] = v : 决策第i件物品, 重量为w时  的最大价值v
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, w+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < w+1; j++ {
			dp[i][j] = -1
		}
	}
	// init state
	dp[0][0] = 0
	if weight[0] < w {
		dp[0][weight[0]] = value[0]
	}
	// 动态规划
	for i := 1; i < n; i++ {
		// 不选择物品 i
		for j := 0; j < w+1; j++ {
			if dp[i-1][j] > 0 {
				dp[i][j] = dp[i-1][j]
			}
		}
		// 选择物品 i
		for j := 1; j < w-weight[i]; j++ {
			if dp[i-1][j] >= 0 {
				dp[i][j+weight[i]] = max(dp[i][j+weight[i]], dp[i][j]+value[i])
			}
		}
	}
	// find max
	var maxValue int
	for i := 0; i <= w; i++ {
		maxValue = max(maxValue, dp[n-1][i])
	}
	return maxValue
}
