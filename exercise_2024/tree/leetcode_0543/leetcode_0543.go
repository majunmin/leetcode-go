package leetcode_0543

import . "github.com/majunmin/leetcode-go/common"

func diameterOfBinaryTree(root *TreeNode) int {
	d, _ := diameterOfBinaryTreeHelper(root)
	return d
}

// @return 最大直径 最大半径
func diameterOfBinaryTreeHelper(root *TreeNode) (int, int) {
	// terminate
	if root == nil {
		return 0, 0
	}
	lmd, lmr := diameterOfBinaryTreeHelper(root.Left)
	rmd, rmr := diameterOfBinaryTreeHelper(root.Right)
	return max(lmd, rmd, lmr+rmr), max(lmr, rmr) + 1
}
