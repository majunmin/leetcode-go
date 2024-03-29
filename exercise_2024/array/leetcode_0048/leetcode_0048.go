package leetcode_0048

// https://leetcode.cn/problems/rotate-image/?envType=study-plan-v2&envId=top-100-liked
func rotate(matrix [][]int) {
	// param check
	if len(matrix) == 0 || len(matrix[0]) == 0 || len(matrix) != len(matrix[0]) {
		panic("invalid param")
	}
	m := len(matrix)
	// 两次翻转 对称翻转
	for i := 1; i < m; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 中轴线翻转
	for i := 0; i < m/2; i++ {
		for j := 0; j < m; j++ {
			matrix[j][i], matrix[j][m-i] = matrix[j][m-i], matrix[j][i]
		}
	}
}
