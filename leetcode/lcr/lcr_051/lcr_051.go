package lcr_051

import . "github.com/majunmin/leetcode-go/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	// dfs
	if root == nil {
		return 0
	}
	res, _ := maxPathSumHelp(root)
	return res
}

// @return 最大路径和 + 最大单条路径和
func maxPathSumHelp(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	lres, lPathSum := maxPathSumHelp(node.Left)
	rres, rPathSum := maxPathSumHelp(node.Right)
	return max(lres, rres, lPathSum+rPathSum+node.Val),
		max(0, max(lPathSum, rPathSum)+node.Val)
}
