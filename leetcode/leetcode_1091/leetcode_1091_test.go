package leetcode_1091

import (
	"fmt"
	"testing"
)

func TestShortPath(t *testing.T) {
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 1, 0},
		{0, 0, 0, 1, 1, 0},
		{1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 0},
	}))
}
