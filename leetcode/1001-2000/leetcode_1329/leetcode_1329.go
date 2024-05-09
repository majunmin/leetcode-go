package leetcode_1329

// https://leetcode.cn/problems/sort-the-matrix-diagonally/
func diagonalSort(mat [][]int) [][]int {
	// param check
	var (
		m, n = len(mat), len(mat[0])
	)
	if m < 2 || n < 2 {
		return mat
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 排序,
			decr := 1
			cur := mat[i][j]
			//插入排序
			for i-decr >= 0 && j-decr >= 0 {
				if cur >= mat[i-decr][j-decr] {
					break
				}
				mat[i-decr+1][j-decr+1] = mat[i-decr][j-decr]
				decr++
			}
			mat[i-decr+1][j-decr+1] = cur
		}
	}
	return mat
}
