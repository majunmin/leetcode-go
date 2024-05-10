package leetcode_0036

// https://leetcode.cn/problems/valid-sudoku/?envType=study-plan-v2&envId=top-interview-150
func isValidSudoku(board [][]byte) bool {
	var (
		rows   = make(map[int]map[byte]bool)
		cols   = make(map[int]map[byte]bool)
		blocks = make(map[int]map[byte]bool)
	)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			if rows[i] == nil {
				rows[i] = make(map[byte]bool)
			}
			if !valid(board[i][j], rows[i]) {
				return false
			}
			if cols[j] == nil {
				cols[j] = make(map[byte]bool)
			}
			if !valid(board[i][j], cols[j]) {
				return false
			}
			blockIdx := i/3*3 + j/3
			if blocks[blockIdx] == nil {
				blocks[blockIdx] = make(map[byte]bool)
			}
			if !valid(board[i][j], blocks[blockIdx]) {
				return false
			}
		}
	}
	return true
}

func valid(item byte, visited map[byte]bool) bool {

	if visited[item] {
		return false
	} else {
		visited[item] = true
	}
	return true
}

func isValidSudoku2(board [][]byte) bool {
	var (
		rows   = make([][9]int, 9)
		cols   = make([][9]int, 9)
		blocks = make([][3][9]int, 3)
	)
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}
			item := board[i][j] - '0' - 1
			rows[i][item]++
			cols[j][item]++
			blocks[i/3][j/3][item]++
			// check
			if rows[i][item] > 1 || cols[j][item] > 1 || blocks[i/3][j/3][item] > 1 {
				return false
			}
		}
	}
	return true
}
