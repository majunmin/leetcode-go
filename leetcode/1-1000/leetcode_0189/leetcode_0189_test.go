// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-26 23:23
package leetcode_0189

import (
	"fmt"
	"testing"
)

func Test_Rotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(nums, 2)

	fmt.Println(nums)
}
