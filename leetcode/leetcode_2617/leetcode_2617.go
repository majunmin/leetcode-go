package leetcode_2617

import "math"

// https://leetcode.cn/problems/minimum-number-of-visited-cells-in-a-grid/
func minimumVisitedCells(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, i int }
	colStack := make([][]pair, 0) //每列的单调栈
	rowSt := []pair{}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {

		}
	}
	
}

type point struct {
	x, y int
}

// 会报错超出内存限制, 再次基础上进行优化
func dpBfsSolution(grid [][]int) int {
	// 看着 求最小最大值， 而不是具体的路径,第一感觉就是动态规划的题目
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	//dp := util.Make2Slice[int](m, n)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	// init state
	dp[0][0] = 1
	// bfs 进行遍历
	var (
		queue = make([]point, 1)
	)
	queue = append(queue, point{0, 0})
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			x, y := queue[i].x, queue[i].y
			// 探索行
			end := min(n-1, y+grid[x][y])
			for k := y + 1; k <= end; k++ {
				dp[x][k] = min(dp[x][k], dp[x][y]+1)
				queue = append(queue, point{x, k})
			}
			// 探索列
			end = min(m-1, x+grid[x][y])
			for k := x + 1; k <= end; k++ {
				dp[k][y] = min(dp[k][y], dp[x][y]+1)
				queue = append(queue, point{k, y})
			}
		}
		queue = queue[size:]
	}
	if dp[m-1][n-1] == math.MaxInt {
		return -1
	}
	return dp[m-1][n-1]
}
