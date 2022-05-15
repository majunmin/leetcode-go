package leetcode_0098

import (
	"fmt"
	. "github.com/majunmin/leetcode-go/common"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	node1 := &TreeNode{Val: 0}
	node2 := &TreeNode{Val: -1}
	//node3 := &TreeNode{Val: 4}
	//node4 := &TreeNode{Val: 3}
	//node5 := &TreeNode{Val: 6}
	node1.Left = node2
	//node1.Right = node3
	//node3.Left = node4
	//node3.Right = node5

	fmt.Println(isValidBST(node1))

}
