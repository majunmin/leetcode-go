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

/**
当前点击的是`未挖出的地雷M`，我们将其值改为 `X` **即可**.
当前点击的是`未挖出的空方块E`，我们需要统计它周围相邻的方块里地雷的数量 `cnt`（即 `M` 的数量).
     如果 `cnt` == 0, 即执行规则 2，此时需要将其改为 `B`.且递归地处理周围的八个未挖出的方块，递归终止条件即为规则 44，没有更多方块可被揭露的时候。
     否则执行规则 3 将其修改为数字**即可**.

*/

// https://leetcode.cn/problems/minesweeper/description/
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
	visited := make([]bool, len(board[0])*len(board))
	// 放入队列后 立即加入到  visited 里面，防止重复入queue,造成OOM
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

			// count E 周围  M  的数量
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
				if !isValidPoint(board, newRowIdx, newColIdx) ||
					board[newRowIdx][newColIdx] != 'E' || visited[newRowIdx*len(board[0])+newColIdx] {
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
func dfsSolution(board [][]byte, rowIdx, colIdx int) {
	// terminated
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
		if isValidPoint(board, newRowIdx, newColIdx) && board[newRowIdx][newColIdx] == 'E' {
			dfsSolution(board, newRowIdx, newColIdx)
		}
	}
}

func isValidPoint(board [][]byte, newRowIdx int, newColIdx int) bool {
	return newRowIdx >= 0 && newRowIdx < len(board) && newColIdx >= 0 && newColIdx < len(board[0])
}

func isDigit(b byte) bool {
	return b >= '1' && b <= '9'
}
