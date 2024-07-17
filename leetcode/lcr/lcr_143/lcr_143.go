package lcr_143

import (
	. "github.com/majunmin/leetcode-go/common"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// param check
	return (A != nil && B != nil) && (hasSubStructure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func hasSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return hasSubStructure(A.Left, B.Left) && hasSubStructure(A.Right, B.Right)
}
