package leetcode_3122

import "math"

// https://leetcode.cn/problems/minimum-number-of-operations-to-satisfy-conditions/
func minimumOperations(grid [][]int) int {
	return (grid)
}

// solution1 自底向上 解决方式
func solution1(grid [][]int) int {
	var (
		m, n = len(grid), len(grid[0])
		// cnts[i][j] grid中第i列的 值为 j 的数出现次数.
		cnts = make([][10]int, n)
		// memo[i][j] 表示考虑前 i 列，且第 i+1 列都变成 j 时的最大保留不变元素个数.
		memo = make([][]int, n)
	)
	// init state
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, 11)
		for j := 0; j < 11; j++ {
			memo[i][j] = -1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnts[j][grid[i][j]]++
		}
	}
	// j 是一个不在 0-9 之间的数
	return m*n - dfs1(n-1, 10, cnts, memo)
}

// prev 上一列数字
func dfs1(i int, prev int, cnts [][10]int, memo [][]int) int {
	// 1. terminate
	if i < 0 {
		return 0
	}
	// 2. memo
	if memo[i][prev] != -1 { //之前计算过
		return memo[i][prev]
	}
	var res int
	// 3. 暴力枚举每一个数
	for k := 0; k < 10; k++ {
		// 是同一列， skip
		if k != prev {
			// 4. dfs
			res = max(res, dfs1(i-1, k, cnts, memo)+cnts[i][k])
		}
	}
	memo[i][prev] = res // 记忆化
	return res
}

// solution2 自顶向下 解决
func solution2(grid [][]int) int {
	var (
		m, n = len(grid), len(grid[0])
		// 统计每一列中,每个数字出现的次数.
		cnts = make([][10]int, n)
	)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnts[j][grid[i][j]]++
		}
	}

	return dfs2(grid, 0, -1, cnts)
}

func dfs2(grid [][]int, cur int, prev int, cnts [][10]int) int {
	if cur == len(grid[0]) {
		return 0
	}
	var (
		minChange = math.MaxInt
	)
	for num, cnt := range cnts[cur] {
		if num == prev {
			continue
		}
		minChange = min(minChange, dfs2(grid, cur+1, cur, cnts)-cnt)
	}
	return minChange
}
