package leetcode_1091

var directions = [][]int{
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

// https://leetcode.cn/problems/shortest-path-in-binary-matrix/
func shortestPathBinaryMatrix(grid [][]int) int {
	// param check
	rowSize, colSize := len(grid), len(grid[0])
	if rowSize == 0 || colSize == 0 {
		return 0
	}
	if grid[0][0] == 1 {
		return -1
	}

	var queue []int
	visited := make([][]int, 0, rowSize)
	for i := 0; i < colSize; i++ {
		visited = append(visited, make([]int, colSize))
	}
	// init
	queue = append(queue, 0)
	minDistance := 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			row, col := queue[i]/rowSize, queue[i]%colSize
			visited[row][col] = 1

			// terminate
			if row == rowSize-1 && col == colSize-1 {
				return minDistance
			}
			// 探索路径

			for _, dir := range directions {
				newRow, newCol := row+dir[0], col+dir[1]
				if newRow < 0 || newRow >= rowSize || newCol < 0 || newCol >= colSize ||
					grid[newRow][newCol] == 1 || visited[newRow][newCol] == 1 {
					continue
				}
				queue = append(queue, newRow*rowSize+newCol)
			}
		}
		minDistance++
		queue = queue[size:]
	}
	// 走不到终点
	return -1
}
