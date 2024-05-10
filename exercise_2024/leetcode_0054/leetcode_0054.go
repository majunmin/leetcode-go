package leetcode_0054

// https://leetcode.cn/problems/spiral-matrix/?envType=study-plan-v2&envId=top-interview-150
func spiralOrder(matrix [][]int) []int {
	// 模拟
	var (
		left, right = 0, len(matrix[0]) - 1
		up, down    = 0, len(matrix) - 1

		x, y   int
		result = make([]int, 0, len(matrix)*len(matrix[0]))
	)
	for {
		// 向右→
		for y = left; y <= right; y++ {
			result = append(result, matrix[up][y])
		}
		up++
		if up > down {
			break
		}
		// 🔽向下
		for x = up; x <= down; x++ {
			result = append(result, matrix[x][right])
		}
		right--
		if left > right {
			break
		}
		// 向左◀️
		for y = right; y >= left; y-- {
			result = append(result, matrix[down][y])
		}
		down--
		if up > down {
			break
		}
		// 向上🔼
		for x = down; x >= up; x-- {
			result = append(result, matrix[x][left])
		}
		left++
		if left > right {
			break
		}
	}
	return result
}

var (
	directions = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)

func spiralOrder2(matrix [][]int) []int {
	// param check
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	// 模拟
	var (
		rowSize, colSize = len(matrix), len(matrix[0])
		result           []int
		visited          = make([][]bool, rowSize)
		dirIdx           int
		total            = rowSize * colSize
		x, y             int
	)
	for i := 0; i < rowSize; i++ {
		visited[i] = make([]bool, colSize)
	}
	for i := 0; i < total; i++ {
		result = append(result, matrix[x][y])
		visited[x][y] = true
		nextx, nexty := x+directions[dirIdx][0], y+directions[dirIdx][1]
		for nextx < 0 || nextx >= rowSize || nexty < 0 || nexty >= colSize || visited[nextx][nexty] {
			dirIdx = (dirIdx + 1) % 4
			nextx, nexty = x+directions[dirIdx][0], y+directions[dirIdx][1]
			break
		}
		x, y = nextx, nexty
	}
	return result
}
