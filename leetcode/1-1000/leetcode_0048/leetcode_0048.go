package leetcode_0048

// 实现矩阵的转置
// https://leetcode.cn/problems/rotate-image/description/?envType=study-plan-v2&envId=top-100-liked
// 题解： https://leetcode.cn/problems/rotate-image/solutions/526980/xuan-zhuan-tu-xiang-by-leetcode-solution-vu3m/?envType=study-plan-v2&envId=top-100-liked
func rotate(matrix [][]int) {
	// 1. 上下翻转 矩阵
	// 2. 沿主对角线翻转矩阵
	rowSize, colSize := len(matrix), len(matrix[0])
	for i := 0; i < rowSize/2; i++ {
		for j := 0; j < colSize; j++ {
			matrix[i][j] = matrix[rowSize-i-1][j]
		}
	}
	// 2.
	for i := 1; i < rowSize; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j] = matrix[j][i]
		}
	}
}
