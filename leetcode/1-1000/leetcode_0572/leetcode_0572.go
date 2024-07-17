package leetcode_0572

import . "github.com/majunmin/leetcode-go/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//https://leetcode.cn/problems/subtree-of-another-tree/
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}

	return (root.Val == subRoot.Val && isSame(root.Left, subRoot.Left) && isSame(root.Right, subRoot.Right)) ||
		isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSame(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && isSame(t1.Left, t2.Left) && isSame(t1.Right, t2.Right)
}
