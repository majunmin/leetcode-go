package leetcode_0226

import . "github.com/majunmin/leetcode-go/common"

func invertTree(root *TreeNode) *TreeNode {
	// terminate
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
