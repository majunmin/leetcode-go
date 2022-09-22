package offer_036

import (
	"fmt"
	"testing"

	. "github.com/majunmin/leetcode-go/offer"
)

func Test_TreeToDoubleList(t *testing.T) {
	root := &TreeNode{Val: 1}
	l := &TreeNode{Val: 2}
	r := &TreeNode{Val: 3}
	root.Left, root.Right = l, r

	l.Left, l.Right = &TreeNode{Val: 4}, &TreeNode{Val: 5}
	head := treeToDoublyList(root)
	fmt.Println(head)
	fmt.Println("len")
}
