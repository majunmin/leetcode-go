package leetcode_0289

var (
	directions = [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
	}
)

// https://leetcode.cn/problems/game-of-life/?envType=study-plan-v2&envId=top-interview-150
func gameOfLife(board [][]int) {
	// 需要有状态标记 出来 该位置原来是 0 还是1
	// 0 -> 0 : 0
	// 1 -> 1 : 1
	// 0 -> 1 : -1
	// 1 -> 0 : 2
	// <= 0 的 的位置原来是 死的
	// >  0 的 的位置原来是 活的
	var (
		m, n = len(board), len(board[0])
	)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var liveCnt, deadCnt int
			for _, d := range directions {
				newx, newy := i+d[0], j+d[1]
				if newx < 0 || newx >= m || newy < 0 || newy >= n {
					if board[newx][newy] > 0 {
						liveCnt++
					} else {
						deadCnt++
					}
				}
			}
			if board[i][j] == 0 {
				if liveCnt == 3 {
					board[i][j] = -1
				}
			} else {
				if deadCnt != 2 && deadCnt != 3 {
					board[i][j] = 2
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == -1 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}
