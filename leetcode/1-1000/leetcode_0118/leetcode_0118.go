package leetcode_0118

// https://leetcode.cn/problems/pascals-triangle/?envType=study-plan-v2&envId=top-100-liked
func generate(numRows int) [][]int {
	// param check
	if numRows <= 0 {
		panic("invalid param")
	}
	result := make([][]int, numRows)
	// init
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0], result[i][i] = 1, 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}
