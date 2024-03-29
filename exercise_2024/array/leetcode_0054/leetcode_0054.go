package leetcode_0054

// https://leetcode.cn/problems/spiral-matrix/?envType=study-plan-v2&envId=top-100-liked
func spiralOrder(matrix [][]int) []int {
	// param check
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	//
	m, n := len(matrix), len(matrix[0])
	total := m * n
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	// 模拟 螺旋遍历
	row, col := 0, 0
	var dirIdx int
	var result []int
	for i := 0; i < total; i++ {
		result = append(result, matrix[row][col])
		visited[row][col] = true
		nextRow, nextCol := row+dirs[dirIdx][0], col+dirs[dirIdx][1]
		if nextRow < 0 || nextRow >= m || nextCol < 0 || nextCol >= n || visited[nextRow][nextCol] {
			dirIdx = (dirIdx + 1) % len(dirs)
		}
		row += dirs[dirIdx][0]
		col += dirs[dirIdx][1]
	}
	return result
}
