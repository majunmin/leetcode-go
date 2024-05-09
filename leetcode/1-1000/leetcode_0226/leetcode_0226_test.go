package leetcode_0226

import (
	"fmt"
	"testing"

	"github.com/majunmin/leetcode-go/common"
)

func Test_InvertTree(t *testing.T) {
	node1 := &common.TreeNode{Val: 1}
	node2 := &common.TreeNode{Val: 2}
	node3 := &common.TreeNode{Val: 3}
	node4 := &common.TreeNode{Val: 4}
	node5 := &common.TreeNode{Val: 5}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5

	invertTree(node1)

	fmt.Printf("%v", node1)
}
