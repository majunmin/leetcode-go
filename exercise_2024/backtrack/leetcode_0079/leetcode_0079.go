package leetcode_0079

var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}
)

// https://leetcode.cn/problems/word-search/?envType=study-plan-v2&envId=top-100-liked
func exist(board [][]byte, word string) bool {
	// param check
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	if len(word) == 0 {
		return true
	}
	// process
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != word[0] {
				continue
			}

			visited[i][j] = true
			if backtrace(board, i, j, visited, word, 1) {
				return true
			}
			visited[i][j] = false
		}
	}
	return false
}

func backtrace(board [][]byte, i int, j int, visited [][]bool, word string, idx int) bool {
	// terminate
	if idx == len(word) {
		return true
	}
	// for choice in choiceList
	for _, dir := range directions {
		newi, newj := i+dir[0], j+dir[1]
		if newi < 0 || newi >= len(board) || newj < 0 || newj >= len(board[0]) || visited[newi][newj] {
			continue
		}
		if word[idx] != board[newi][newj] {
			continue
		}
		visited[i][j] = true
		// do it
		if backtrace(board, newi, newj, visited, word, idx+1) {
			return true
		}
		visited[i][j] = false
	}
	return false
}
