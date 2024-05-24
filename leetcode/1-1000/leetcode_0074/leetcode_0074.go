package leetcode_0074

// https://leetcode.cn/problems/search-a-2d-matrix/?envType=study-plan-v2&envId=top-interview-150
func searchMatrix(matrix [][]int, target int) bool {
	// param check
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	left, right := -1, m*n
	for left+1 < right {
		mid := left + (right-left)>>1
		if matrix[mid/n][mid%n] >= target {
			right = mid
		} else {
			left = mid
		}
	}
	return matrix[right/n][right%n] == target
}
