package leetcode_0120

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}
