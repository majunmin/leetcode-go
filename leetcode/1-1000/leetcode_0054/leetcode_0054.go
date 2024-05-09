package leetcode_0054

var (
	directions = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)

// https://leetcode.cn/problems/spiral-matrix/?envType=study-plan-v2&envId=top-100-liked
func spiralOrder(matrix [][]int) []int {
	return layerSolution(matrix)
}

func layerSolution(matrix [][]int) []int {
	// param check
	if len(matrix) == 0 {
		return nil
	}
	var result []int
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	for top < bottom && left < right {
		for i := left; i < right; i++ {
			result = append(result, matrix[top][i])
		}
		for i := top; i < bottom; i++ {
			result = append(result, matrix[i][right])
		}
		for i := right; i > left; i-- {
			result = append(result, matrix[bottom][i])
		}
		for i := bottom; i > top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		right--
		top++
		bottom--
	}
	//
	if top == bottom {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
	}
	if left == right {
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][left])
		}
	}
	return result
}

// 模拟运动解法
func moniSolution(matrix [][]int) []int {
	idx := 0
	rowSize, colSize := len(matrix), len(matrix[0])
	total := rowSize * colSize
	result := make([]int, 0, total)
	row, col := 0, 0
	visited := make([][]bool, 0, rowSize)
	for i := 0; i < rowSize; i++ {
		visited = append(visited, make([]bool, colSize))
	}
	for i := 0; i < total; i++ {
		result = append(result, matrix[row][col])
		visited[row][col] = true
		dir := directions[idx]
		nextRow, nextCol := row+dir[0], col+dir[1]
		if nextRow < 0 || nextRow >= rowSize || nextCol < 0 || nextCol >= colSize ||
			visited[nextRow][nextCol] {
			idx = (idx + 1) % len(directions)
		}
		row, col = row+directions[idx][0], col+directions[idx][1]
	}
	return result
}
