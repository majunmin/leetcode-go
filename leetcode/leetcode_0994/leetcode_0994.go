package leetcode_0994

var (
	directions = [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
)

type point struct {
	x, y int
}

// https://leetcode.cn/problems/rotting-oranges/?envType=study-plan-v2&envId=top-100-liked
func orangesRotting(grid [][]int) int {
	// param check
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return 0
	}
	rowSize, colSize := len(grid), len(grid[0])
	// bfs solution
	queue := make([]point, 0)
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, point{i, j})
			}
		}
	}
	var minutes int
	// process
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			for _, dir := range directions {
				newx, newy := item.x+dir[0], item.y+dir[1]
				if newx < 0 || newx >= rowSize || newy < 0 || newy > colSize || grid[newx][newy] != 1 {
					continue
				}
				grid[newx][newy] = 2
				queue = append(queue, point{newx, newy})
			}
		}
		queue = queue[size:]
	}

	// 统计 剩余新鲜橘子数
	var freshCnt int
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if grid[i][j] == 1 {
				freshCnt++
			}
		}
	}
	if freshCnt > 0 {
		return -1
	}
	return minutes
}
