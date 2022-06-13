package leetcode_0529

var (
	directions = [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}
)

//https://leetcode.cn/problems/minesweeper/description/
func updateBoard(board [][]byte, click []int) [][]byte {
	// param check
	if len(board) == 0 || len(board[0]) == 0 || len(click) != 2 {
		return nil
	}
	bfsSolution(board, Point{x: click[0], y: click[1]})
	return board
}

type Point struct {
	x, y int
}

func bfsSolution(board [][]byte, point Point) {

	if board[point.x][point.y] == 'M' {
		board[point.x][point.y] = 'X'
		return
	}

	queue := make([]Point, 0)
	queue = append(queue, point)
	visited := make([]bool, len(board)*len(board[0]))
	visited[point.x*len(board[0])+point.y] = true
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			p := queue[i]

			if isDigit(board[p.x][p.y]) {
				continue
			}

			if board[p.x][p.y] == 'M' {
				continue
			}

			count := 0
			for _, dir := range directions {
				newRowIdx, newColIdx := p.x+dir[0], p.y+dir[1]
				if !isValidPoint(board, newRowIdx, newColIdx) || board[newRowIdx][newColIdx] != 'M' {
					continue
				}
				count++
			}
			if count > 0 {
				board[p.x][p.y] = '0' + byte(count)
				continue
			}
			board[p.x][p.y] = 'B'
			for _, dir := range directions {
				newRowIdx, newColIdx := p.x+dir[0], p.y+dir[1]
				if !isValidPoint(board, newRowIdx, newColIdx) || visited[newRowIdx*len(board[0])+newColIdx] {
					continue
				}
				visited[newRowIdx*len(board[0])+newColIdx] = true
				queue = append(queue, Point{newRowIdx, newColIdx})
			}
		}
		queue = queue[size:]
	}
}

// dfs solution
func dfsSolution(board [][]byte, rowIdx, colIdx int, visited []bool) {
	// terminated
	if visited[rowIdx*len(board[0])+colIdx] {
		return
	}
	visited[rowIdx*len(board[0])+colIdx] = true

	if board[rowIdx][colIdx] == 'M' {
		board[rowIdx][colIdx] = 'X'
		return
	}

	if isDigit(board[rowIdx][colIdx]) {
		return
	}

	// process cur point
	count := 0
	for _, dir := range directions {
		newRowIdx, newColIdx := rowIdx+dir[0], colIdx+dir[1]
		if isValidPoint(board, newRowIdx, newColIdx) {
			if board[newRowIdx][newColIdx] == 'M' {
				count++
			}
		}
	}
	if count > 0 {
		board[rowIdx][colIdx] = '0' + byte(count)
		return
	}
	board[rowIdx][colIdx] = 'B'

	for _, dir := range directions {
		newRowIdx, newColIdx := rowIdx+dir[0], colIdx+dir[1]
		if isValidPoint(board, newRowIdx, newColIdx) {
			dfsSolution(board, newRowIdx, newColIdx, visited)
		}
	}
}

func isValidPoint(board [][]byte, newRowIdx int, newColIdx int) bool {
	return newRowIdx >= 0 && newRowIdx < len(board) && newColIdx >= 0 && newColIdx < len(board[0])
}

func isDigit(b byte) bool {
	return b >= '1' && b <= '9'
}
