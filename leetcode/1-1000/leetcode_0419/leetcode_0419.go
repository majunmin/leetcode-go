package leetcode_0419

// https://leetcode.cn/problems/battleships-in-a-board/?envType=daily-question&envId=2024-06-11
func countBattleships(board [][]byte) int {
	// bfs
	var (
		m, n   = len(board), len(board[0])
		result int
	)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != 'X' {
				continue
			}
			if (i > 0 && board[i-1][j] == 'X') || (j > 0 && board[i][j-1] == 'X') {
				continue
			}
			result++
		}
	}
	return result
}
