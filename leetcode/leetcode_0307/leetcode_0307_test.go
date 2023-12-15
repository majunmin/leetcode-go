package leetcode_0307

import (
	"fmt"
	"testing"
)

func Test_Case(t *testing.T) {
	na := Constructor([]int{1, 3, 5})
	fmt.Println(na.SumRange(0, 2))
	na.Update(1, 2)
	fmt.Println(na.SumRange(0, 2))
}
