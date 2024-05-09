package leetcode_0543

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/diameter-of-binary-tree/?envType=study-plan-v2&envId=top-100-liked
func diameterOfBinaryTree(root *TreeNode) int {
	// param check
	if root == nil {
		return 0
	}
	res, _ := diameterOfBinaryTreeHelper(root)
	return res
}

// curMaxDiameter  maxRadius
func diameterOfBinaryTreeHelper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lmax, lMaxPath := diameterOfBinaryTreeHelper(root.Left)
	rmax, rMaxPath := diameterOfBinaryTreeHelper(root.Right)
	return max(lmax, rmax, lMaxPath+rMaxPath), max(lMaxPath, rMaxPath) + 1
}
