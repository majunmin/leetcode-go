package offer_028

import (
	"fmt"
	"github.com/majunmin/leetcode-go/offer"
	"testing"
)

func Test_IsSymmetric(t *testing.T) {
	root := &offer.TreeNode{Val: 1}
	node1 := &offer.TreeNode{Val: 2}
	node2 := &offer.TreeNode{Val: 2}
	node3 := &offer.TreeNode{Val: 3}
	node4 := &offer.TreeNode{Val: 4}
	node5 := &offer.TreeNode{Val: 4}
	node6 := &offer.TreeNode{Val: 3}

	root.Left, root.Right = node1, node2
	node1.Left, node1.Right = node3, node4
	node2.Left, node2.Right = node5, node6

	fmt.Println(isSymmetric(root))
}
