package _024_06_23

func minimumArea(grid [][]int) int {
	var (
		m, n            = len(grid), len(grid[0])
		l, r, top, down = n - 1, 0, m - 1, 0
	)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			l = min(l, j)
			r = max(r, j)
			top = min(top, i)
			down = max(down, i)
		}
	}
	return (r - l + 1) * (down - top + 1)
}
