// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-17 21:02
package leetcode_0283

import (
	"fmt"
	"testing"
)

func Test_MoveZeros(t *testing.T) {
	nums := []int{0, 1, 0, 3, 2}
	moveZeroes(nums)

	fmt.Println(nums)

	nums = []int{1, 0, 3, 0, 0, 2}
	moveZeroes(nums)

	fmt.Println(nums)
}
