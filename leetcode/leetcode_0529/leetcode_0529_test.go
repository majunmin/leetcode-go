package leetcode_0529

import (
	"fmt"
	"testing"
)

func Test_UpdateBoard(t *testing.T) {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}

	updateBoard(board, []int{3, 0})
	fmt.Println(board)
}
