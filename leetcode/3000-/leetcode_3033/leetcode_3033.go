package leetcode_3033

// https://leetcode.cn/problems/modify-the-matrix/
func modifiedMatrix(matrix [][]int) [][]int {
	// param check

	var (
		m, n   = len(matrix), len(matrix[0])
		ans    = make([][]int, m)
		colMax = make([]int, n)
	)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = matrix[i][j]
			colMax[j] = max(colMax[j], matrix[i][j])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ans[i][j] == -1 {
				ans[i][j] = colMax[j]
			}
		}
	}
	return ans
}

func modifiedMatrix2(matrix [][]int) [][]int {
	// param check

	var (
		m, n = len(matrix), len(matrix[0])
		ans  = make([][]int, m)
	)
	// 遍历列, 找到每一列的最大值·
	for i := 0; i < n; i++ {
		mx := 0
		// 遍历行， find max col value
		for j := 0; j < m; j++ {
			mx = max(mx, matrix[j][i])
		}
		// 遍历行
		for j := 0; j < m; j++ {
			if len(ans[i]) == 0 {
				ans[i] = make([]int, n)
			}
			ans[j][i] = matrix[j][i]
			if ans[j][i] == -1 {
				ans[j][i] = mx
			}
		}
	}
	return ans
}
