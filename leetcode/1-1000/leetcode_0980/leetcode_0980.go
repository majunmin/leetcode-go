package leetcode_0980

var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

// https://leetcode.cn/problems/unique-paths-iii/
func uniquePathsIII(grid [][]int) int {
	return dsfSolution(grid)

}

func dsfSolution(grid [][]int) int {
	if len(grid) == 0 || len(grid[1]) == 0 {
		return -1
	}

	var (
		sr, sc int
		todo   int
	)
	rowSize, colSize := len(grid), len(grid[0])
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if grid[i][j] == 1 {
				sr, sc = i, j
				continue
			}
			if grid[i][j] != -1 {
				todo++
			}
		}
	}
	return dfs(grid, sr, sc, todo)
}

func dfs(grid [][]int, row, col int, todo int) int {
	todo--
	rowSize, colSize := len(grid), len(grid[0])
	if row < 0 || row >= rowSize || col < 0 || col >= colSize {
		return 0
	}

	if grid[row][col] == 2 && todo == 0 {
		return 1
	}
	if grid[row][col] == -1 || grid[row][col] == 2 {
		return 0
	}
	var result int
	tmp := grid[row][col]
	grid[row][col] = -1
	for _, dir := range directions {
		result += dfs(grid, row+dir[0], col+dir[1], todo)
	}
	grid[row][col] = tmp
	return result
}

func backTraceSolution1(grid [][]int) int {
	rowSize, colSize := len(grid), len(grid[0])

	if rowSize < 1 || colSize < 1 {
		return 0
	}
	var (
		sr, sc = 0, 0
		todo   int
	)
	// find start and end
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if grid[i][j] == 1 {
				sr, sc = i, j
				continue
			}
			if grid[i][j] != -1 {
				todo++
			}
		}
	}

	visisted := make([]int, rowSize*colSize)
	return backTrace(grid, sr, sc, visisted, todo)
}

func backTrace(grid [][]int, row, col int, visited []int, todo int) int {

	rowSize, colSize := len(grid), len(grid[0])
	if visited[row*colSize+col] == 1 {
		return 0
	}
	if grid[row][col] == -1 {
		return 0
	}
	if grid[row][col] == 2 {
		if todo == 0 {
			return 1
		}
		return 0
	}
	todo--
	visited[row*colSize+col] = 1
	// revert
	// backtrace
	var result int
	for _, dir := range directions {
		if row+dir[0] < 0 || row+dir[0] >= rowSize || col+dir[1] < 0 || col+dir[1] >= colSize {
			continue
		}
		result += backTrace(grid, row+dir[0], col+dir[1], visited, todo)
	}
	visited[row*colSize+col] = 0
	return result
}
