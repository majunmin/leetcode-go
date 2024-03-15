package leetcode_1261

import (
	"fmt"
	"testing"

	. "github.com/majunmin/leetcode-go/common"
)

var (
	tree1 = &TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val: -1,
			Left: &TreeNode{
				Val: -1,
				Left: &TreeNode{
					Val:   -1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   -1,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: -1,
			Left: &TreeNode{
				Val: -1,
				Left: &TreeNode{
					Val:   -1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   -1,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   -1,
				Left:  nil,
				Right: nil,
			},
		},
	}
)

func TestConstructor(t *testing.T) {
	elesFinder := Constructor(tree1)
	b1 := elesFinder.Find(8)
	fmt.Print(b1)
	b2 := elesFinder.Find(3)
	fmt.Print(b2)
}
