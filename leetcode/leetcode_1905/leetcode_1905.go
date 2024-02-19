package leetcode_1905

var (
	directions = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)

// https: //leetcode.cn/problems/count-sub-islands/
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	// param check
	if len(grid1) == 0 || len(grid2) == 0 || len(grid1[0]) == 0 || len(grid2[0]) == 0 {
		return 0
	}
	if len(grid1) != len(grid2) || len(grid1[0]) != len(grid2[0]) {
		return 0
	}
	// process
	var result int
	rowSize, colSize := len(grid1), len(grid1[0])
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if grid2[i][j] == 0 {
				continue
			}
			// dfs
			if dfs(grid1, grid2, rowSize, colSize, i, j) {
				result++
			}
		}
	}
	return result
}

// 判断是否是子岛屿
// 如果不是子岛屿  grid1[i][j] != grid2[i][j](== 1)
func dfs(grid1 [][]int, grid2 [][]int, rowSize int, colSize int, i int, j int) bool {
	// 陆地打平
	grid2[i][j] = 0
	res := grid1[i][j] == 1
	for _, dir := range directions {
		newx, newy := i+dir[0], j+dir[1]
		if newx < 0 || newx >= rowSize || newy < 0 || newy >= colSize || grid2[newx][newy] == 0 {
			continue
		}
		res = dfs(grid1, grid2, rowSize, colSize, newx, newy) && res
	}
	return res
}
