package offer_027

import (
	. "github.com/majunmin/leetcode-go/offer"
)

//https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/?envType=study-plan&id=lcof
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left = mirrorTree(root.Left)
	}
	if root.Right != nil {
		root.Right = mirrorTree(root.Right)
	}

	// swap
	root.Left, root.Right = root.Right, root.Left
	return root
}
