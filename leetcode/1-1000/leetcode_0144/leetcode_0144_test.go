package leetcode_0144

import (
	"fmt"
	"testing"

	. "github.com/majunmin/leetcode-go/common"
)

func Test_preorderTraversal(t *testing.T) {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5

	res := preorderTraversal(node1)
	fmt.Println(res)
}
