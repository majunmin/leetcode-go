package offer_068

import . "github.com/majunmin/leetcode-go/offer"

// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/?envType=study-plan&id=lcof
// 二叉搜索树
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// param check
	if root == nil || p == nil || q == nil {
		return nil
	}

	// terminate
	if p.Val == root.Val || q.Val == root.Val {
		return root
	}

	if p.Val < root.Val && q.Val > root.Val {
		return root
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}
