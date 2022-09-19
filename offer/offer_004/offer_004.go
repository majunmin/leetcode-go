package offer_004

func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 矩阵为空
	if len(matrix) == 0 {
		return false
	}
	// query
	row, col := len(matrix)-1, 0
	for row >= 0 && col <= len(matrix[0]) { // check row and col is valid
		if target > matrix[row][col] {
			col += 1
		} else if target < matrix[row][col] {
			row -= 1
		} else {
			return true
		}
	}
	return false
}
