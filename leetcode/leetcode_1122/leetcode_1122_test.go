package leetcode_1122

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	arr1 := []int{26, 21, 11, 20, 50, 34, 1, 18}
	arr2 := []int{21, 11, 26, 20}
	relativeSortArray(arr1, arr2)
	fmt.Println(arr1)
}
