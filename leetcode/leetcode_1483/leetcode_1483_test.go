package leetcode_1483

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	treeAncestor := Constructor(7, []int{-1, 0, 0, 1, 1, 2, 2})
	fmt.Println(treeAncestor.GetKthAncestor(3, 1))
	fmt.Println(treeAncestor.GetKthAncestor(5, 2))
	fmt.Println(treeAncestor.GetKthAncestor(4, 3))
}
