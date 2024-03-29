package leetcode_0236

import (
	. "github.com/majunmin/leetcode-go/common"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// param check
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	ln := lowestCommonAncestor(root.Left, p, q)
	rn := lowestCommonAncestor(root.Right, p, q)
	if ln == nil {
		return rn
	}
	if rn == nil {
		return ln
	}
	return root
}
