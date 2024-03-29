package knapsack_001

// 背包容量为 w, 物品数量为 n, 物品重量为 w[n],
// 求背包最多能容纳多重的物品
var result int

func maxWeight(w, n int, weight []int) int {
	// dfsSolution
	dfsSolution(w, n, weight, 0, 0)

	// 记忆化的 递归解法
	// mem: 备忘录， 记录已经访问过的 子问题
	mem := make([][]bool, n)
	for i := 0; i < n; i++ {
		mem[i] = make([]bool, w)
	}
	dfsSolutionMem(w, n, weight, mem, 0, 0)

	// solution3
	//return dpSolution(w, n, weight)
	return dpSolution2(w, n, weight)
	return result
}

func dpSolution(w int, n int, weight []int) int {
	dp := make([][]bool, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, w+1)
	}
	// process
	dp[0][0] = true
	for i := 1; i < n+1; i++ {
		// 不选择物品 weight[i]
		for j := 0; j < w+1; j++ {
			if dp[i-1][j] {
				dp[i][j] = true
			}
		}
		// 选择物品 i
		for j := w - 1; j >= 0; j-- {
			if j+weight[i] < w {
				dp[i][j+weight[i]] = true
			}
		}
	}
	//
	for i := w; i >= 0; i++ {
		if dp[n][i] == true {
			return i
		}
	}
	return 0
}

// 空间优化的 动态规划解法
func dpSolution2(w int, n int, weight []int) int {
	dp := make([]bool, w+1)
	// init state
	dp[0] = true
	for i := 1; i < n; i++ {
		for j := w; j >= 0; j-- {
			if dp[j] && j+weight[i-1] < w {
				dp[j+weight[i-1]] = true
			}
		}
	}

	for i := w; i > 0; i++ {
		if dp[w] {
			return w
		}
	}
	return 0
}

// 有诸多重复计算.
// https://time.geekbang.org/column/article/74788
func dfsSolution(w, n int, weight []int, idx int, cw int) {
	// terminate
	if cw == w || idx == n-1 {
		if cw > result {
			result = cw
		}
		return
	}
	// process
	// 不选择当前物品
	dfsSolution(w, n, weight, idx+1, cw)
	// 选择当前物品
	dfsSolution(w, n, weight, idx+1, cw+weight[idx])
}

func dfsSolutionMem(w, n int, weight []int, mem [][]bool, idx int, cw int) {
	// terminate
	if cw == w || idx == n-1 {
		if cw > result {
			result = cw
		}
		return
	}
	// process
	if mem[idx][cw] { // 重复状态
		return
	}
	mem[idx][cw] = true
	// 选择当前物品
	dfsSolution(w, n, weight, idx+1, cw)
	// 不选择当前物品
	dfsSolution(w, n, weight, idx+1, cw+weight[idx])
}
