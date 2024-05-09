package leetcode_0102

import (
	"fmt"
	"testing"

	. "github.com/majunmin/leetcode-go/common"
)

func Test_LevelOrder(t *testing.T) {
	root := &TreeNode{Val: 1}
	rl := &TreeNode{Val: 3}
	rr := &TreeNode{Val: 5}
	rrl := &TreeNode{Val: 7}
	rrr := &TreeNode{Val: 15}
	root.Left = rl
	root.Right = rr
	rr.Left = rrl
	rr.Right = rrr

	result := levelOrder(root)
	fmt.Println(result)
}
