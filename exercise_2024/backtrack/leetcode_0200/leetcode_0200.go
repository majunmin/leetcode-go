package leetcode_0200

var (
	directions = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)

func numIslands(grid [][]byte) int {
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		// param invalid
		return 0
	}
	rowSize, colSize := len(grid), len(grid[0])
	var result int
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if grid[i][j] == '0' {
				continue
			}
			bfs(grid, rowSize, colSize, i, j)
			result++
		}
	}
	return result
}

type point struct {
	x, y int
}

func bfs(grid [][]byte, rowSize int, colSize int, i int, j int) {
	queue := make([]point, 0)
	queue = append(queue, point{i, j})
	for len(queue) > 0 {
		size := len(queue)
		for k := 0; k < size; k++ {
			item := queue[k]
			if grid[item.x][item.y] == '0' {
				continue
			}
			grid[item.x][item.y] = '0'
			for _, dir := range directions {
				newx, newy := item.x+dir[0], item.y+dir[1]
				if newx < 0 || newx >= rowSize || newy < 0 || newy >= colSize {
					continue
				}
				queue = append(queue, point{x: newx, y: newy})
			}
		}
		queue = queue[size:]
	}
}

func dfs(grid [][]byte, rowSize int, colSize int, i int, j int) {
	if i < 0 || i >= rowSize || j < 0 || j > colSize {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	for _, dir := range directions {
		dfs(grid, rowSize, colSize, i+dir[0], j+dir[1])
	}
}
