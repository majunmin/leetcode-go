package leetcode_0085

// https://leetcode.cn/problems/maximal-rectangle/
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var (
		m, n = len(matrix), len(matrix[0])
	)
	
}
