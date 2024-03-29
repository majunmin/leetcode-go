package leetocde_0235

import . "github.com/majunmin/leetcode-go/common"

// 公共祖先的条件.
// node == p == q
// node.Left contains p, node.Right contains q
// node == p,   node.Right contains q
// node == q,   node.Left contains q
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}
