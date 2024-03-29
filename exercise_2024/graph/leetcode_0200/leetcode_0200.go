package leetcode_0200

var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}
)

func numIslands(grid [][]byte) int {
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	var cnt int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt++
				// 将陆地打平
				bfs(grid, i, j)
			}
		}
	}
	return cnt
}

type point struct {
	x, y int
}

func bfs(grid [][]byte, x int, y int) {
	queue := make([]point, 0)
	queue = append(queue, point{x, y})
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			grid[item.x][item.y] = '0'
			for _, dir := range directions {
				newx, newy := item.x+dir[0], item.y+dir[1]
				if newx < 0 || newx >= len(grid) || newy < 0 || newy >= len(grid[0]) || grid[newx][newy] == '0' {
					continue
				}
				queue = append(queue, point{newx, newy})
			}
		}
		queue = queue[size:]
	}
}

func dfs(grid [][]byte, i int, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	for _, dir := range directions {
		newx, newy := i+dir[0], j+dir[1]
		dfs(grid, newx, newy)
	}
}

func numIslandsUfSolution(grid [][]byte) int {
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var (
		m, n   = len(grid), len(grid[0])
		spaces int
		uf     = NewUnionFind(m * n)
		dirs   = [][]int{
			{-1, 0},
			{0, -1},
		}
	)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				spaces++
				continue
			}
			for _, dir := range dirs {
				newx, newy := i+dir[0], j+dir[1]
				if newx < 0 || newy < 0 || newx >= m || newy >= n || grid[newx][newy] == '0' {
					continue
				}
				uf.union(getIndex(i, j, n), getIndex(newx, newy, n))
			}
		}
	}
	return uf.cnt - spaces
}

func getIndex(i int, j int, colSize int) int {
	return i*colSize + j
}

type UnionFind struct {
	cnt    int
	parent []int
}

func NewUnionFind(cnt int) *UnionFind {
	parent := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		parent[i] = i
	}
	return &UnionFind{
		cnt:    cnt,
		parent: parent,
	}
}

func (uf *UnionFind) Len() int {
	return uf.cnt
}
func (uf *UnionFind) find(x int) int {
	for uf.parent[x] != x {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

func (uf *UnionFind) union(x, y int) {
	rootx, rooty := uf.find(x), uf.find(y)
	if rootx == rooty {
		return
	}
	uf.parent[rootx] = rooty
	uf.cnt--
}
