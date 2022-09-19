package offer_012

//https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/

var directions = [][]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

type Node struct {
	x int
	y int
}

//
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 {
		return false
	}
	colSize := len(board[0])

	for i := range board {
		for j := range board[i] {
			idx := 0
			if board[i][j] != word[idx] {
				continue
			}

			var queue []Node
			queue = append(queue, Node{i, j})

			visited := make([]bool, len(board)*colSize)
			visited[i*colSize+j] = true
			for len(queue) > 0 {
				size := len(queue)
				idx++
				if idx == len(word) {
					return true
				}
				for m := 0; m < size; m++ {
					node := queue[m]
					for _, direction := range directions {
						for dx, dy := range direction {
							newI, newJ := dx+node.x, node.y+dy
							if 0 > newI || newI >= len(board) || 0 > newJ || newJ >= colSize || visited[newI*colSize+newJ] {
								continue
							}
							if board[newI][newJ] != word[idx] {
								continue
							}
							visited[newI*colSize+newJ] = true
							queue = append(queue, Node{newI, newJ})
						}
					}
				}
				queue = queue[size:]
			}
		}
	}

	return false

}

// dfsSolution
func dfsSolution(board [][]byte, word string) bool {
	// param check
	if len(word) == 0 {
		return false
	}
	visited := make([]bool, len(board)*len(board[0]))
	for i := range board {
		for j := range board[i] {
			if dfs(board, visited, i, j, []byte(word), 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, visited []bool, i int, j int, word []byte, idx int) bool {
	colSize := len(board[0])
	// 剪枝
	if word[idx] != board[i][j] {
		return false
	}
	// terminate
	if idx+1 == len(word) {
		return true
	}

	visited[i*colSize+j] = true
	for _, direction := range directions {
		newI, newJ := i+direction[0], j+direction[1]
		// 剪枝
		if newI >= 0 && newI < len(board) && newJ >= 0 && newJ < colSize && !visited[newI*colSize+newJ] {
			if dfs(board, visited, newI, newJ, word, idx+1) {
				return true
			}
		}

	}
	visited[i*colSize+j] = false
	return false
}
