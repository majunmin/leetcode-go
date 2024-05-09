package leetcode_0615

var (
	direction = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

// https://leetcode.cn/problems/shortest-bridge/
func shortestBridge(grid [][]int) int {
	// dfs + bfs
	// 1. 找到第一个 岛屿, 将 1 反转为  2, 从而区分 两个岛屿
	// 2. 对找到的岛屿进行 BFS遍历, 直到找到 1
	// 3. 遍历的层数 就是  最短距离

	// param check
	queue := make([]Point, 0)
	var one Point
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				one = Point{x: i, y: j}
				break
			}
		}
	}
	// dfs反转 1 -> 2
	dfs(grid, one, &queue)
	// bfs 遍历
	var minDistance int
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			p := queue[i]
			for _, dir := range direction {
				newx, newy := p.x+dir[0], p.y+dir[1]
				// valid
				if newx < 0 || newx > len(grid) || newy < 0 || newy > len(grid[0]) {
					continue
				}
				// 已经被访问过,skip
				if grid[newx][newy] == 2 {
					continue
				}

				// terminate
				if grid[newx][newy] == 1 {
					return minDistance
				}

				// 标记已经遍历过的 点
				grid[newx][newy] = 2
				queue = append(queue, Point{x: newx, y: newy})
			}

		}
		minDistance++
		queue = queue[n:]
	}

	return minDistance
}

func dfs(grid [][]int, one Point, queue *[]Point) {
	if one.x < 0 || one.x >= len(grid) || one.y < 0 || one.y > len(grid[0]) {
		return
	}

	if grid[one.x][one.y] != 1 {
		return
	}
	*queue = append(*queue, one)

	grid[one.x][one.y] = 2
	for _, dir := range direction {
		dfs(grid, Point{x: one.x + dir[0], y: one.y + dir[1]}, queue)
	}
}

type Point struct {
	x int
	y int
}
