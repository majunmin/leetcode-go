package leetcode_0073

var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}
)

// https://leetcode.cn/problems/set-matrix-zeroes/?envType=study-plan-v2&envId=top-100-liked
func setZeroes(matrix [][]int) {
	solution3(matrix)
}

func solution3(matrix [][]int) {
	rowSize, colSize := len(matrix), len(matrix[0])
	row0, col0 := false, false
	for i := 0; i < rowSize; i++ {
		if matrix[i][0] == 0 {
			col0 = true
			break
		}
	}
	for i := 0; i < colSize; i++ {
		if matrix[0][i] == 0 {
			row0 = true
			break
		}
	}
	//
	for i := 1; i < rowSize; i++ {
		for j := 1; j < colSize; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < rowSize; i++ {
		for j := 1; j < colSize; j++ {

			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for i := 0; i < colSize; i++ {
			matrix[0][i] = 0
		}
	}
	if col0 {
		for i := 0; i < rowSize; i++ {
			matrix[i][0] = 0
		}
	}
}

// 使用一个flag 标记, 节省了空间， 但是不好理解
func solution4(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	col0 := false
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}

func solution2(matrix [][]int) {
	rowFlag, colFlag := make([]bool, len(matrix)), make([]bool, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rowFlag[i] = true
				colFlag[j] = true
			}
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			if rowFlag[i] || colFlag[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func solution1(matrix [][]int) {
	queue := make([][]int, 0)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	rowSize, colSize := len(matrix), len(matrix[0])
	for _, point := range queue {
		for _, dir := range directions {
			nx, ny := point[0], point[1]
			for ; nx >= 0 && nx < rowSize && ny >= 0 && ny < colSize; nx, ny = point[0]+dir[0], point[1]+dir[1] {
				matrix[nx][ny] = 0
			}
		}
	}
}
