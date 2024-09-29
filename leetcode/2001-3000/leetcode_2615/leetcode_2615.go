package leetcode_2615

// https://leetcode.cn/problems/sum-of-distances/description/
//func distance(nums []int) []int64 {
//
//}

var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var (
		m, n    = len(grid), len(grid[0])
		zeroCnt int
		uf      = newUnionFind(m * n)
	)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				zeroCnt++
				continue
			}
			// result++
			// bfs(grid, i, j)
			if i > 0 && grid[i-1][j] == '1' {
				uf.union(i*n+j, (i-1)*n+j)
			}
			if j > 0 && grid[i][j-1] == '1' {
				uf.union(i*n+j, i*n+j-1)
			}
		}
	}

	// return result
	return uf.cnt - zeroCnt
}

func dfs(grid [][]byte, row, col int) {
	// terminate
	grid[row][col] = '0'
	for _, dir := range directions {
		newRow, newCol := row+dir[0], col+dir[1]
		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) || grid[newRow][newCol] == '0' {
			continue
		}
		dfs(grid, newRow, newCol)
	}
}

type point struct {
	x, y int
}

func bfs(grid [][]byte, row, col int) {
	var (
		queue = make([]point, 0)
	)
	grid[row][col] = '0'
	queue = append(queue, point{row, col})
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			// process
			for _, dir := range directions {
				newRow, newCol := item.x+dir[0], item.y+dir[1]
				if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) || grid[newRow][newCol] == '0' {
					continue
				}
				grid[newRow][newCol] = '0'
				queue = append(queue, point{x: newRow, y: newCol})
			}
		}
		queue = queue[size:]
	}
}

type unionFind struct {
	p    []int
	rank []int
	cnt  int
}

func newUnionFind(size int) *unionFind {
	p := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		p[i] = i
		rank[i] = 1
	}
	return &unionFind{
		p:    p,
		rank: rank,
		cnt:  size,
	}
}

func (uf *unionFind) union(i, j int) {
	rootx, rooty := uf.find(i), uf.find(j)
	if rootx == rooty {
		return
	}
	// merge
	if uf.rank[rootx] == uf.rank[rooty] {
		uf.p[rootx] = rooty
		uf.rank[rooty]++
	} else if uf.rank[rootx] < uf.rank[rooty] {
		uf.p[rootx] = rooty
	} else {
		uf.p[rooty] = rootx
	}
	uf.cnt--
}

func (uf *unionFind) find(i int) int {
	if uf.p[i] != i {
		uf.p[i] = uf.find(uf.p[i])
	}
	return uf.p[i]
}
