package offer_012

//https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/
var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

func exist(board [][]byte, word string) bool {

	// param check
	rowSize := len(board)
	if rowSize == 0 {
		return false
	}
	colSize := len(board[0])
	if colSize == 0 {
		return false
	}

	visited := make([][]bool, rowSize)
	for i := 0; i < rowSize; i++ {
		visited[i] = make([]bool, colSize)
	}
	//
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			if backTrace(board, visited, i, j, &[]byte{}, []byte(word), 0) {
				return true
			}
		}
	}
	return false

}

func backTrace(board [][]byte, visited [][]bool, rowIdx int, colIdx int, path *[]byte, word []byte, idx int) bool {
	*path = append(*path, board[rowIdx][colIdx])
	//剪枝
	if word[idx] != board[rowIdx][colIdx] {
		return false
	}
	if idx == len(word) {
		return true
	}
	visited[rowIdx][colIdx] = true

	for _, dir := range directions {
		newRow, newCol := rowIdx+dir[0], colIdx+dir[1]
		if 0 <= newRow && newRow < len(board) && 0 <= newCol && newCol < len(board[0]) && !visited[newRow][newCol] {
			if backTrace(board, visited, newRow, newCol, path, word, idx+1) {
				return true
			}
		}
	}
	visited[rowIdx][colIdx] = false
	*path = (*path)[:len(*path)-1]
	return false
}
