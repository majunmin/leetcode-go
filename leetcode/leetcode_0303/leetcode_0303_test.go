package leetcode_0303

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	numArr := Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	fmt.Println(numArr.SumRange(0, 0))
	fmt.Println(numArr.SumRange(0, 1))
	fmt.Println(numArr.SumRange(0, 3))
	fmt.Println(numArr.SumRange(2, 3))
	fmt.Println(numArr.SumRange(9, 9))
}
