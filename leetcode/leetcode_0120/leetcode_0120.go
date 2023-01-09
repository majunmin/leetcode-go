package leetcode_0120

// https://leetcode.cn/problems/triangle/description/
func minimumTotal(triangle [][]int) int {
	return solution1(triangle)

}

func solution1(triangle [][]int) int {
	n := len(triangle)
	// param check

	// init arr
	temp := make([][]int, n)
	for i := 0; i < n; i++ {
		temp[i] = make([]int, n)
	}

	temp[0][0] = triangle[0][0]
	// 动态规划
	for i := 1; i < n; i++ {
		temp[i][0] = triangle[i][0] + temp[i-1][0]
		for j := 1; j < len(triangle[i])-1; j++ {
			temp[i][j] = triangle[i][j] + minInt(temp[i-1][j], temp[i-1][j-1])
		}
		temp[i][len(triangle[i])-1] = triangle[i][len(triangle[i])-1] + temp[i-1][len(triangle[i])-2]
	}

	minEle := temp[n-1][0]
	for i := 1; i < n; i++ {
		minEle = minInt(minEle, temp[n-1][i])
	}
	return minEle
}

// 自底向上 解决, 修改了 原始结构
func solution3(triangle [][]int) int {
	// param check
	if len(triangle) == 0 {
		return -1
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	//
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += minInt(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
