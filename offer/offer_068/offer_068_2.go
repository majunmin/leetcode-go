package offer_068

import . "github.com/majunmin/leetcode-go/offer"

// https://leetcode.cn/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/?envType=study-plan&id=lcof
// 二叉树的后序遍历
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	//  terminate
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	//
	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)
	if leftNode != nil && rightNode != nil {
		return root
	}
	if leftNode != nil {
		return leftNode
	}
	return rightNode
}
