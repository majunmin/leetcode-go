package leetcode_0994

var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}
)

// https://leetcode.cn/problems/rotting-oranges/?envType=study-plan-v2&envId=top-100-liked
func orangesRotting(grid [][]int) int {
	// bfs 遍历
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	//
	var (
		m, n     = len(grid), len(grid[0])
		queue    = make([]point, 0)
		freshCnt int
		seconds  int
	)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, point{i, j})
			} else if grid[i][j] == 1 {
				freshCnt++
			}
		}
	}

	for len(queue) > 0 {
		if freshCnt == 0 {
			seconds++
			break
		}
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			grid[item.x][item.y] = 0
			for _, dir := range directions {
				newx, newy := item.x+dir[0], item.y+dir[1]
				if newx < 0 || newx >= m || newy < 0 || newy >= n || grid[newx][newy] == 0 || grid[newx][newy] == 2 {
					continue
				}
				// 此时  grid[newx][newy] == 1
				grid[newx][newy] = 2
				freshCnt--
				queue = append(queue, point{newx, newy})
			}
		}
		queue = queue[size:]
	}
	return seconds
}

type point struct {
	x, y int
}
