package leetcode_1644

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree-ii/description/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var (
		findp, findq          bool
		lowestCommonAncestor1 func(*TreeNode, *TreeNode, *TreeNode) *TreeNode
	)
	//
	lowestCommonAncestor1 = func(node *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node == p {
			findp = true
		}
		if node == q {
			findq = true
		}
		l := lowestCommonAncestor1(node.Left, p, q)
		r := lowestCommonAncestor1(node.Right, p, q)
		if node == p || node == q || (l != nil && r != nil) {
			return node
		}
		if r != nil {
			return r
		}
		return l
	}
	ans := lowestCommonAncestor1(root.Left, p, q)
	if findp && findq {
		return ans
	}
	return nil
}
