// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-26 23:16
package leetcode_0026

import (
	"fmt"
	"testing"
)

func Test_RemoveDuplicateNums(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1, 2, 3, 4, 4, 5, 5, 6, 7, 7, 8}))
	fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates([]int{1, 1, 1}))
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 2, 3}))
}
