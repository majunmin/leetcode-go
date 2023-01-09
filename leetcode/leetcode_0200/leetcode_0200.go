package leetcode_0200

var (
	diretions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

// https://leetcode.cn/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	// paramCheck
	var queue []int
	var count int
	rowSize := len(grid)
	colSize := len(grid[0])
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if grid[i][j] == '0' {
				continue
			}
			grid[i][j] = '0'

			count++
			queue = append(queue, i*colSize+j)
			for len(queue) > 0 {
				size := len(queue)
				for k := 0; k < size; k++ {
					m, n := queue[k]/colSize, queue[k]%colSize
					for _, dir := range diretions {
						newRow, newCol := m+dir[0], n+dir[1]
						if 0 <= newRow && newRow < rowSize && 0 <= newCol && newCol < colSize && grid[newRow][newCol] == '1' {
							grid[newRow][newCol] = '0'
							queue = append(queue, newRow*colSize+newCol)
						}
					}
				}
				queue = queue[size:]
			}
		}
	}
	return count
}

func uniFindSolution(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rowSize, colSize := len(grid), len(grid[0])
	uf := NewUnionFind(rowSize * colSize)
	for i := 0; i < rowSize; i++ {
		for j := 0; j < rowSize; j++ {

		}
	}
	return uf.Count()
}

func dsfSolution(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	count := 0
	// 每个点的 标记 i* colSize + j
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				continue
			}
			// 岛屿数量 + 1
			count++
			// process
			dsf(grid, i, j)
		}
	}
	return count
}

func dsf(grid [][]byte, row int, col int) {
	// 打平
	grid[row][col] = '0'
	for _, dir := range diretions {
		newRow, newCol := row+dir[0], col+dir[1]
		if 0 <= newRow && newRow < len(grid) && 0 <= newCol && newCol < len(grid[0]) {
			// 遇到 1 就打平
			if grid[newRow][newCol] == '1' {
				dsf(grid, newRow, newCol)
			}
		}
	}

}
