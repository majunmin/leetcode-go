package leetcode_0130

var (
	directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

// 1. 使用 额外的空间 表示外层的'O'
// 2. 优化空间:使用特殊的标记 标记外层 'O'
func solve(board [][]byte) {
	// 反向思考 🤔
	var (
		m, n = len(board), len(board[0])
		temp = make([][]byte, m)
	)
	for i := 0; i < m; i++ {
		temp[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			temp[i][j] = board[i][j]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && i < m-1 && j > 0 && j < n-1 {
				continue
			}
			if board[i][j] == 'X' {
				continue
			}
			dfs(board, i, j)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if temp[i][j] == 'X' {
				continue
			}
			if board[i][j] == 'X' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	board[i][j] = 'X'
	for _, dir := range directions {
		newi, newj := i+dir[0], j+dir[1]
		if newi < 0 || newi >= len(board) || newj < 0 || newj >= len(board[0]) || board[newi][newj] == 'X' {
			continue
		}
		dfs(board, newi, newj)
	}
}
