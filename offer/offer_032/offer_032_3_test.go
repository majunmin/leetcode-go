package offer_032

import (
	"fmt"
	"testing"
)

func Test_ReverseLevel(t *testing.T) {
	root := &TreeNode{Val: 0}
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n6 := &TreeNode{Val: 6}
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4
	n2.Left = n5
	n2.Right = n6
	fmt.Println(levelOrder3(root))
}
