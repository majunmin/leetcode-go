// Created By: junmin.ma
// Description: <description>
// Date: 2022-01-11 22:04
package leetcode_1036

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	blocked := [][]int{}
	source := []int{0, 1}
	target := []int{2, 2}
	fmt.Println(isEscapePossible(blocked, source, target))
}
