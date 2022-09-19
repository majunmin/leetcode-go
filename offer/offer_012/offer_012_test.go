package offer_012

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCED"))
	//fmt.Println(exist([][]byte{{'a'}, {'a'}, {'a'}}, "aaaa"))
}
