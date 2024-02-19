package leetcode_0236

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	// terminate
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// process
	lson := lowestCommonAncestor(root.Left, p, q)
	rson := lowestCommonAncestor(root.Right, p, q)
	if lson != nil || rson != nil {
		return root
	}
	if lson == nil {
		return rson
	}
	return lson
}
