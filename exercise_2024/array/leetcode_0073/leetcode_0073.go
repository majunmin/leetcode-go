package leetcode_0073

// https://leetcode.cn/problems/set-matrix-zeroes/?envType=study-plan-v2&envId=top-100-liked
func setZeroes(matrix [][]int) {
	// param check
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	rowFlag, colFlag := make([]bool, m), make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowFlag[i], colFlag[j] = true, true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowFlag[i] || colFlag[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func solution1(matrix [][]int) {
	// param check
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	var rowFlag, colFlag bool
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colFlag = true
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			rowFlag = true
		}
	}
	//
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}
	// set zero
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if rowFlag {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if colFlag {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
