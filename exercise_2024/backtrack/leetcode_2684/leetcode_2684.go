package leetcode_2684

import "math"

var (
	directions = [][]int{
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	revDirs = [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
	}
)

// https://leetcode.cn/problems/maximum-number-of-moves-in-a-grid/
func maxMoves(grid [][]int) int {
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	return dpSolution(grid)
}

func dpSolution(grid [][]int) int {
	dp := make([][]int, len(grid))
	var result int
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
		dp[i][0] = 1
	}
	for j := 1; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			for _, dir := range revDirs {
				px, py := i+dir[0], j+dir[1]
				if px < 0 || py < 0 || px >= len(grid) || py >= len(grid[0]) ||
					dp[px][py] == 0 || grid[i][j] <= grid[px][py] {
					continue
				}
				dp[i][j] = max(dp[i][j], dp[px][py]+1)
			}
			result = max(result, dp[i][j])
		}
	}
	return result - 1
}package leetcode_2684

import "math"

var (
	directions = [][]int{
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	revDirs = [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
	}
)

// https://leetcode.cn/problems/maximum-number-of-moves-in-a-grid/
func maxMoves(grid [][]int) int {
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	return dpSolution(grid)
}

func dpSolution(grid [][]int) int {
	dp := make([][]int, len(grid))
	var result int
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
		dp[i][0] = 1
	}
	for j := 1; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			for _, dir := range revDirs {
				px, py := i+dir[0], j+dir[1]
				if px < 0 || py < 0 || px >= len(grid) || py >= len(grid[0]) ||
					dp[px][py] == 0 || grid[i][j] <= grid[px][py] {
					continue
				}
				dp[i][j] = max(dp[i][j], dp[px][py]+1)
			}
			result = max(result, dp[i][j])
		}
	}
	return result - 1
}

type point struct {
	x, y int
}

// 2. bfs 解法
func bfsSolution(grid [][]int) int {
	queue := make([]point, 0)
	m, n := len(grid), len(grid[0])
	var cnt int
	for i := 0; i < m; i++ {
		p := point{i, 0}
		queue = append(queue, p)
	}
	for len(queue) > 0 {
		cnt++
		pointSet := make(map[point]bool)
		size := len(queue)
		for i := 0; i < size; i++ {
			p := queue[i]
			for _, dir := range directions {
				newx, newy := p.x+dir[0], p.y+dir[1]
				if newx < 0 || newy < 0 || newx >= m || newy >= n || grid[p.x][p.y] >= grid[newx][newy] || pointSet[point{newx, newy}] {
					continue
				}
				newP := point{newx, newy}
				pointSet[newP] = true
				queue = append(queue, newP)
			}
		}
		queue = queue[size:]
	}
	return cnt - 1
}

// 1. dfs 超时
func dfsSolution(grid [][]int) int {
	var maxVal int
	for i := 0; i < len(grid); i++ {
		maxVal = max(maxVal, dfs(grid, math.MinInt, i, 0))
	}
	return maxVal
}

func dfs(grid [][]int, prevVal, row, col int) int {
	// terminate
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || prevVal >= grid[row][col] {
		return 0
	}
	if col == len(grid[0])-1 {
		return 1
	}
	// for choice in choiceList
	var cnt int
	for _, dir := range directions {
		nx, ny := row+dir[0], col+dir[1]
		cnt = max(cnt, dfs(grid, grid[row][col], nx, ny))
	}
	return cnt + 1
}


type point struct {
	x, y int
}

// 2. bfs 解法
func bfsSolution(grid [][]int) int {
	queue := make([]point, 0)
	m, n := len(grid), len(grid[0])
	var cnt int
	for i := 0; i < m; i++ {
		p := point{i, 0}
		queue = append(queue, p)
	}
	for len(queue) > 0 {
		cnt++
		pointSet := make(map[point]bool)
		size := len(queue)
		for i := 0; i < size; i++ {
			p := queue[i]
			for _, dir := range directions {
				newx, newy := p.x+dir[0], p.y+dir[1]
				if newx < 0 || newy < 0 || newx >= m || newy >= n || grid[p.x][p.y] >= grid[newx][newy] || pointSet[point{newx, newy}] {
					continue
				}
				newP := point{newx, newy}
				pointSet[newP] = true
				queue = append(queue, newP)
			}
		}
		queue = queue[size:]
	}
	return cnt - 1
}

// 1. dfs 超时
func dfsSolution(grid [][]int) int {
	var maxVal int
	for i := 0; i < len(grid); i++ {
		maxVal = max(maxVal, dfs(grid, math.MinInt, i, 0))
	}
	return maxVal
}

func dfs(grid [][]int, prevVal, row, col int) int {
	// terminate
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || prevVal >= grid[row][col] {
		return 0
	}
	if col == len(grid[0])-1 {
		return 1
	}
	// for choice in choiceList
	var cnt int
	for _, dir := range directions {
		nx, ny := row+dir[0], col+dir[1]
		cnt = max(cnt, dfs(grid, grid[row][col], nx, ny))
	}
	return cnt + 1
}
