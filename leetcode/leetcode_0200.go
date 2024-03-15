package leetcode

var (
	directions = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)

type point struct {
	x, y int
}

func numIslands(grid [][]byte) int {
	rowSize, colSize := len(grid), len(grid[0])
	if rowSize <= 0 || colSize <= 0 {
		return 0
	}
	var cnt int
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if grid[i][j] == '0' {
				continue
			}
			cnt++
			//bfs(grid, i, j, rowSize, colSize)
			dfs200(grid, i, j, rowSize, colSize)
		}
	}
	return cnt
}

func dfs200(grid [][]byte, i int, j int, rowSize, colSize int) {
	// terminate
	if i < 0 || i > rowSize || j < 0 || j > colSize || grid[i][j] == '0' {
		return
	}
	// process
	grid[i][j] = '0'
	for _, dir := range directions {
		dfs200(grid, i+dir[0], j+dir[1], rowSize, colSize)
	}
}

func bfs(grid [][]byte, i int, j int, rowSize int, colSize int) {
	// 将岛屿打平 bfs
	// visited： 由于访问过的节点就置为 0 了,所以不用记录访问过的节点
	queue := make([]point, 0)
	queue = append(queue, point{i, j})
	grid[i][j] = '0'
	for len(queue) > 0 {
		size := len(queue)
		for k := 0; k < size; k++ {
			item := queue[k]
			for _, dir := range directions {
				newx, newy := item.x+dir[0], item.y+dir[1]
				if newx >= 0 && newx < rowSize && newy >= 0 && newy < colSize && grid[newx][newy] == '1' {
					queue = append(queue, point{newx, newy})
					// process point
					grid[newx][newy] = '0'
				}
			}
		}
		queue = queue[size:]
	}
}
